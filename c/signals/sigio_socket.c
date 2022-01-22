#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <sys/socket.h>
#include <unistd.h>
#include <fcntl.h>
#include <errno.h>
#include <signal.h>


#define SERVER_PORT 9876
#define SERVER_BACKLOG 10
#define CLIENT_SNOOZE_TIMER_SECS 5
#define BUFFER_SIZE 128

/**
 * @brief Exit with EXIT_FAILURE if err indicates fatal error
 * 
 * @param err 
 * @param msg 
 */
static inline void check_err_exit(int err, char *msg)
{
    if (err < 0) {
        perror(msg);
        exit(EXIT_FAILURE);
    }
}

/**
 * @brief Set the up server socket in non-blocking mode
 * 
 * @param port 
 * @param backlog 
 * @return int socket descriptor to the server connection
 */
int setup_server(int port, int backlog);

/**
 * @brief Accept new connection on the given non-blocking socket.
 *        Returns -1 if no connection is pending.
 * 
 * @param sockfd 
 * @return int socket descriptor to the client connection.
 */
int accept_connection(int sockfd);

/**
 * @brief signal handler for SIGIO/SIGPOLL signals
 * When the signal is fired, read available data from the client sockets.
 * 
 * @param signum 
 * @param info 
 * @param uap 
 */
void signal_handler(int signum, siginfo_t *info, void *uap);

int setup_server(int port, int backlog)
{
    struct sockaddr_in addr;
    int sockfd;
    int err = -1;

    // create a socket
    sockfd = socket(PF_INET /* protocol family - INET */, 
                    SOCK_STREAM /* socket type - TCP */,
                    0 /* flags */);
    check_err_exit(sockfd, "socket() failed");

    // Set the socket in non-blocking mode.
    int flags = fcntl(sockfd, F_GETFL, 0 /* flags */);
    err = fcntl(sockfd, F_SETFL, flags | O_NONBLOCK);
    check_err_exit(err, "fcntl(F_SETFL) failed");

    // enable reuse of the server port.
    // comes in handy to avoid port in use error when testing
    // in quick intervals. I think the default timeout to
    // cleanup is 20seconds or so.
    int reuse = 1;
    err = setsockopt(sockfd,
               SOL_SOCKET /* option for socket level */,
               SO_REUSEADDR, /* allow reuse of local addr */
               (const void *)&reuse, sizeof(int));
    check_err_exit(err, "setsockopt() failed");
    
    // bind the socket to the given addr:port
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    err = bind(sockfd, (struct sockaddr *)&addr, sizeof(addr));
    check_err_exit(err, "bind() failed");
        
    // set it to listening.
    err = listen(sockfd, backlog);
    check_err_exit(err, "listen() failed");

    printf("Server listening at port %d\n", port);
    return sockfd;
}

int accept_connection(int sockfd)
{
    struct sockaddr_in clientaddr;
    socklen_t addrlen = sizeof(clientaddr);
    
    bzero(&clientaddr, addrlen);
    
    int connection_fd = accept(sockfd, (struct sockaddr *)&clientaddr, &addrlen);
    if (connection_fd == -1) {
        /*
         * If no client connection is waiting, then accept() would
         * return EWOULDBLOCK since the socket is marked non-blocking.
         */
        if (errno != EWOULDBLOCK) {
            check_err_exit(connection_fd, "accept() failed");
        }
    } else {
        printf("Connection from %s:%d, fd=%d\n", 
                    inet_ntoa(clientaddr.sin_addr),clientaddr.sin_port, connection_fd);
    }
    return connection_fd;
}

void signal_handler(int signum, siginfo_t *info, void *uap)
{
    // read until all currently available data is read.
    int bytesread = -1;
    char buf[BUFFER_SIZE] = "\0";

    // si_fd doesn't seem to be populated correctly for some reason
    // signal_handler() is invoked when the client sends the data
    // but I cannot read it without having the right file descriptor
    // info->si_fd comes up as 0

    // I thought info->si_fd will be set to the fd that caused SIGIO
    // to be raised. Apparently, it reports only one fd even when there
    // are multiple fds ready. Since stdin is ready to be read, si_fd
    // comes up with 0, so I couldn't read the data sent by the client

    // why hardcoded to 4? Just playing around.
    // the after stdin, stdout, stderr and the server socket,
    // the descriptors for the incoming connections start at 4.
    // so reading from the first connection to test the behavior
    // when the client sends some data, signal_handler() is invoked.

    // may be this is not the correct pattern of using async I/O
    // we cannot target particular fd.
    // Perhaps, this would work.
    // Keep a pool of descriptors,
    // with all of them in non-blocking mode. 
    // when sigio is fired, go and read from all of them
    // if any of the sockets aren't ready, they would return EAGAIN.
    // this is somewhat equivalent to select/poll but more expensive
    // way of doing what they do.
    int sockfd = 4;//info->si_fd;

    printf("%s: signum=%d, sockfd=%d, si_signo=%d,"
            "si_errno=%d, si_code=0x%x, si_pid=%d\n",
            __func__, signum, sockfd, 
            info->si_signo, info->si_errno, info->si_code, info->si_pid);

    for (;;) {
        bytesread = read(sockfd, (void *)buf, BUFFER_SIZE);
        if (bytesread == -1) {
            if (errno != EAGAIN) {
                check_err_exit(bytesread, "read() failed");
            }
            printf("Read all available data. Waiting for next signal\n");
            break;
        } else if (bytesread == 0) {
            printf("Connection closed by client\n");
            close(sockfd);
            break;
        } else {
            buf[bytesread] = '\0';
            printf("client sent: %s\n", buf);
            write(sockfd, (const void *)buf, bytesread);
        }
    }
}

int main(int argc, char **argv)
{
    int serversock = setup_server(SERVER_PORT, SERVER_BACKLOG);
    int err = -1;

    for (;;) {
        int connection = accept_connection(serversock);
        if (connection == -1) {
            printf("No connection available yet. Going back to sleep\n");
            sleep(CLIENT_SNOOZE_TIMER_SECS);
        } else {
            printf("Setting up connection %d\n", connection);

            // set the current pid as the owner in order to receive the
            // signals
            int pid = getpid();
            err = fcntl(connection, F_SETOWN, pid);
            check_err_exit(err, "fcntl(F_SETOWN) failed");

            // set the client connection to async mode
            // O_ASYNC - tells the kernel to generate SIGIO to the owning
            // pid (set in the previous steps) when I/O is possible on this
            // file descriptor
            int flags = fcntl(connection, F_GETFL, 0);
            fcntl(connection, F_SETFL, flags | O_ASYNC | O_NONBLOCK);

            // set the signal handler
            struct sigaction action;
            action.sa_flags = SA_SIGINFO;
            action.sa_sigaction = signal_handler;
            sigemptyset(&action.sa_mask);
            err = sigaction(SIGIO, &action , NULL);

            // when the client connection is ready to be read, SIGIO will
            // be generated, and the registered handler signal_handler will
            // be invoked.
        }
    }

    // NOT_REACHED
    close(serversock);
    return 0;
}