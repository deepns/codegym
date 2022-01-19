#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <string.h> // for strcmp
#include <sys/socket.h>
#include <sys/epoll.h>
#include <sys/fcntl.h>
#include <unistd.h>
#include <errno.h>

#define BUFFER_SIZE 128
#define SERVER_PORT 11012
#define SERVER_BACKLOG 10
#define MAX_EVENTS 5
#define POLL_TIMEOUT_MS 30000

/*
 * Trying out I/O multiplexing using non-blocking sockets and epoll
 * in edge-triggered mode. Server is a simple echo server that reads
 * a block of data from client and writes the same back to the client.
 * If client says 'bye', server will close the client connection.
 */

/**
 * @brief Set the up server
 *
 * @param port
 * @param backlog
 * @return int socket descriptor to the server connection
 */
int setup_server(int port, int backlog);

/**
 * @brief Accept new connection on the given socket.
 *        Blocks until new connection comes in.
 *
 * @param sockfd
 * @return int socket descriptor to the client connection.
 */
int accept_connection(int sockfd);

/**
 * @brief Handle client connection
 *
 * @param sockfd socket descriptor to the client connection
 * @param closed output param. if not NULL, it will be set with a flag
 *               to indicate whether was closed or not
 */
void handle_connection(int sockfd, int epollfd);

/**
 * @brief Exit with EXIT_FAILURE if err indicates error (-1)
 *
 * @param err
 * @param msg
 */
static inline void check_err(int err, char *msg)
{
    if (err < 0) {
        perror(msg);
        exit(EXIT_FAILURE);
    }
}

/**
 * @brief Set the socket as nonblocking
 * 
 * @param sockfd 
 */
static void set_nonblocking(int sockfd)
{
    int flags = fcntl(sockfd,
                    F_GETFL /* cmd to get value of file status flags */,
                    0 /* optional arg. F_GETFL doesn't take any arg. hence 0 */);                   
    check_err(flags, "fcntl(F_GETFL) failed\n");
    check_err(fcntl(sockfd, F_SETFL, flags | O_NONBLOCK), "fcntl(F_SETFL) failed");
}

int setup_server(int port, int backlog)
{
    struct sockaddr_in addr;
    int sockfd;
    int err = -1;

    // create a socket
    sockfd = socket(PF_INET /* protocol family - INET */,
                    SOCK_STREAM /* socket type - TCP */,
                    0 /* flags */);
    check_err(sockfd, "socket() failed");

    // enable reuse of the server port.
    // comes in handy to avoid port in use error when testing
    // in quick intervals. I think the default timeout to
    // cleanup is 20seconds or so.
    int reuse = 1;
    err = setsockopt(sockfd,
               SOL_SOCKET /* option for socket level */,
               SO_REUSEADDR, /* allow reuse of local addr */
               (const void *)&reuse, sizeof(int));
    check_err(err, "setsockopt() failed");

    set_nonblocking(sockfd);

    // bind the socket to the given addr:port
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    err = bind(sockfd, (struct sockaddr *)&addr, sizeof(addr));
    check_err(err, "bind() failed");

    // set it to listening.
    err = listen(sockfd, backlog);
    check_err(err, "listen() failed");

    printf("Server listening at port %d...Ctrl-C to quit\n", port);
    return sockfd;
}

int accept_connection(int sockfd)
{
    struct sockaddr_in clientaddr;
    bzero(&clientaddr, sizeof(clientaddr));
    socklen_t addrlen = sizeof(clientaddr);

    int connection_fd = accept(sockfd, (struct sockaddr *)&clientaddr, &addrlen);
    if (connection_fd == -1) {
        /*
         * If no client connection is waiting, then accept() would
         * return EWOULDBLOCK since the socket is marked non-blocking.
         */
        if (errno != EWOULDBLOCK) {
            check_err(connection_fd, "accept() failed");
        }
    } else {
        // can also use getpeername() to get this info using
        // the socket descriptor
        printf("Connection from %s:%d\n", inet_ntoa(clientaddr.sin_addr),clientaddr.sin_port);
    }
    return connection_fd;
}

void handle_connection(int sockfd, int epollfd)
{
    char buf[BUFFER_SIZE] = "\0";
    int close_connection = 0;
    ssize_t bytesread = -1; 

    // read until read() returns EAGAIN
    do {
        bytesread = read(sockfd, (void *)buf, BUFFER_SIZE);
        if (bytesread == -1) {
            // if socket is healthy but no data available to read, try again.
            if (errno != EAGAIN) {
                check_err(bytesread, "read() failed");
            }
            break;
        } else if (bytesread == 0) {
            printf("Client closed the connection...\n");
            close_connection = 1;
        } else {
            buf[bytesread] = '\0';
            if (strcmp(buf, "bye") == 0) {
                printf("Closing the connection...bye\n");
                close_connection = 1;
            } else {
                printf("client sent: %s\n", buf);
                write(sockfd, (const void *)buf, bytesread);
            }
        }

        if (close_connection) {
            /*
             * Kernel automatically removes the file descriptor from the
             * epoll interest list when the file descriptor is closed
             * (if no other handle is open to the same descriptor).
             * So EPOLL_CTL_DEL is not really needed here.
             * Explicitly removing it just for experimentation purpose.
             */
            check_err(epoll_ctl(epollfd, EPOLL_CTL_DEL, sockfd, NULL),
                    "epoll_ctl(DEL) failed");
            close(sockfd);
            break;
        }
    } while (1);
}

int main(int argc, char **argv)
{
    int epollfd = -1;
    int n_ready_fds = 0;
    struct epoll_event ev, events[MAX_EVENTS];
    int serversock = setup_server(SERVER_PORT, SERVER_BACKLOG);

    /*
     * epoll_create() used to take a size argument, but has been ignored
     * since Linux 2.6.8.
     * epoll_create1() takes a flag instead. epoll_create1(0) is same as
     * epoll_create() with size ignored. 
     */
    epollfd = epoll_create1(0 /* flags */);
    check_err(epollfd, "epoll_create1() failed");

    /*
     * Register the server socket to the epoll interest list.
     * With select(), we maintain a bit-set of file descriptors to be watched.
     * With poll(), we maintain the list of interested fds in an array.
     * With epoll(), the onus is on the kernel. application simply tells
     * the kernel to add a descriptor to watch using EPOLL_CTL_ADD command.
     * Unlike select()/poll(), epoll doesn't have check all file descriptors
     * each time epoll_wait is called.
     */

    // /proc/sys/fs/epoll/max_user_watches contains the value of total
    // number of file descriptors that a user can register across
    // all epoll instances on the system. On my Ubuntu VM, it is 197857.

    ev.events = EPOLLIN;
    ev.data.fd = serversock;
    check_err(epoll_ctl(epollfd, EPOLL_CTL_ADD, serversock, &ev),
              "epoll_ctl() failed");

    // poll the epoll descriptor for readiness
    // we started with only the server socket. when a client
    // connection comes in, we will add the client socket to
    // the epoll descriptor.
    for (;;) {
        n_ready_fds = epoll_wait(epollfd,
                                events, /* will be populated by epoll_wait */
                                MAX_EVENTS, /* max events to watch for*/
                                POLL_TIMEOUT_MS /* timeout */);
        if (n_ready_fds == 0) {
            printf("epoll_wait timer expired...polling again\n");
            continue;
        } else if (n_ready_fds == -1) {
            check_err(n_ready_fds, "epoll_wait() failed");
        } else {
            // Check which FDs are ready for op
            for (int i = 0; i < n_ready_fds; i++) {
                if (events[i].data.fd == serversock) {
                    int clientsock = accept_connection(serversock);
                    /*
                     * When the poll mode is set to event-triggered,
                     * epoll_wait() will wait until the socket becomes readable.
                     * (i.e. socket received a new write and ready to be read)
                     * Whereas if the poll mode is level-triggered,
                     * epoll_wait() will wait until the socket is readable.
                     * (i.e socket either received a new write or still has
                     * unread data from previous write, so buffer still has
                     * some data to be read).
                     * 
                     * Because epoll_wait edge trigger wakes up only when
                     * there is a new data, setting the client socket to
                     * non-blocking. client connection handler will read
                     * all data until read() returns EAGAIN.
                     */
                    set_nonblocking(clientsock);

                    // add to client socket to epoll watch, with edge trigger.
                    ev.events = EPOLLET | EPOLLIN; // watching for only read.
                    ev.data.fd = clientsock;
                    check_err(epoll_ctl(epollfd, EPOLL_CTL_ADD, clientsock, &ev),
                              "epoll_ctl() failed on clientsock");
                } else {
                    handle_connection(events[i].data.fd, epollfd);
                }
            }
        }
    }

    // NOT_REACHED
    close(serversock);
    return 0;
}
