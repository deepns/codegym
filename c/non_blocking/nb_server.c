#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <sys/socket.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <errno.h>

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
 * @brief Handle client connection
 * 
 * @param sockfd socket descriptor to the client connection
 */
void handle_connection(int sockfd);

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
    int ON = 1;
    err = ioctl(sockfd, FIONBIO, (char *)&ON);
    check_err_exit(err, "ioctl() failed");

    // Another common way to enable non-blocking mode
    // is using fcntl. ioctl predates fcntl, and can
    // be inconsistent between different systems.
    // fcntl is portable. another difference is that
    // ioctl involves only one syscall whereas fcntl
    // involves two syscalls (to get and set).
    // int flags = fcntl(sockfd, F_GETFL, 0);
    // fcntl(sockfd, F_SETFL, flags | O_NONBLOCK);

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
        printf("Connection from %s:%d\n", 
                    inet_ntoa(clientaddr.sin_addr),clientaddr.sin_port);
    }
    return connection_fd;
}

#define BUFFER_SIZE 128

/**
 * @brief Handle client connection
 * 
 * @param sockfd socket descriptor to the client connection
 */
void handle_connection(int sockfd)
{
    char buf[BUFFER_SIZE] = "\0";
    ssize_t bytesread = -1;   
    for (;;) {
        bytesread = read(sockfd, (void *)buf, BUFFER_SIZE);
        if (bytesread == -1) {
            if (errno != EAGAIN) {
                check_err_exit(bytesread, "read() failed");
            }
            printf("Waiting for additional data from client\n");
            sleep(1);
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

#define SERVER_PORT 9699
#define SERVER_BACKLOG 10
#define CLIENT_SNOOZE_TIMER_SECS 5

int main(int argc, char **argv)
{
    int serversock = setup_server(SERVER_PORT, SERVER_BACKLOG);

    for (;;) {
        int connection = accept_connection(serversock);
        // Just getting started with exploring non-blocking sockets
        // Ideally we would want to use poll/epoll to multiplex
        // between the sockets. More on select/poll/epoll in coming days.
        if (connection == -1) {
            printf("No connection available yet. Going back to sleep\n");
            sleep(CLIENT_SNOOZE_TIMER_SECS);
        } else {
            handle_connection(connection);
        }
    }

    // NOT_REACHED
    close(serversock);
    return 0;
}