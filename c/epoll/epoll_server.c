#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <string.h> // for strcmp
#include <sys/socket.h>
#include <sys/epoll.h>
#include <unistd.h>

#define BUFFER_SIZE 128
#define SERVER_PORT 11011
#define SERVER_BACKLOG 10

/*
 * Trying out select to multiplex I/O between multiple sockets.
 * Server is a simple echo server that reads some text from the
 * client and writes it back on the same channel. If client says
 * 'bye', server will close the client connection. We will use
 * 'select' to check for readiness of the server and client sockets
 * and act on them.
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

    // bind the socket to the given addr:port
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    err = bind(sockfd, (struct sockaddr *)&addr, sizeof(addr));
    check_err(err, "bind() failed");

    // set it to listening.
    err = listen(sockfd, backlog);
    check_err(err, "listen() failed");

    printf("Server listening at port %d\n", port);
    return sockfd;
}

int accept_connection(int sockfd)
{
    struct sockaddr_in clientaddr;
    bzero(&clientaddr, sizeof(clientaddr));
    socklen_t addrlen = sizeof(clientaddr);

    int connection_fd = accept(sockfd, (struct sockaddr *)&clientaddr, &addrlen);
    check_err(connection_fd, "accept() failed");

    // can also use getpeername() to get this info using
    // the socket descriptor
    printf("Connection from %s:%d\n", inet_ntoa(clientaddr.sin_addr),clientaddr.sin_port);

    return connection_fd;
}

void handle_connection(int sockfd, int epollfd)
{
    char buf[BUFFER_SIZE] = "\0";
    ssize_t bytesread = read(sockfd, (void *)buf, BUFFER_SIZE);

    buf[bytesread] = '\0';
    if (strcmp(buf, "bye") == 0) {
        printf("Closing the connection...bye\n");
        check_err(epoll_ctl(epollfd, EPOLL_CTL_DEL, sockfd, NULL),
                "epoll_ctl(DEL) failed");
        close(sockfd);
    } else {
        printf("client sent: %s\n", buf);
        write(sockfd, (const void *)buf, bytesread);
    }
}

#define MAX_EVENTS 5
#define POLL_TIMEOUT_MS 30000

int main(int argc, char **argv)
{
    int epollfd = -1;
    int n_ready_fds = 0;
    struct epoll_event ev, events[MAX_EVENTS];
    int serversock = setup_server(SERVER_PORT, SERVER_BACKLOG);

    epollfd = epoll_create1(0 /* flags */);
    check_err(epollfd, "epoll_create1() failed");

    // add the server socket to epoll
    ev.events = EPOLLIN;
    ev.data.fd = serversock;
    check_err(epoll_ctl(epollfd, EPOLL_CTL_ADD, serversock, &ev),
              "epoll_ctl() failed");

    // poll the epoll descriptor for readiness
    // we started with only the server socket. when a client
    // connection comes in, we will add the client socket to
    // the epoll fd.
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
            for (int i=0; i<n_ready_fds; i++) {
                if (events[i].data.fd == serversock) {
                    // accept new connection
                    int clientsock = accept_connection(serversock);

                    // add to client socket to epoll watch
                    // default mode is LEVEL Triggered. leaving it as is for now
                    // will try Edge Triggered separately
                    ev.events = EPOLLIN; // watching for only read.
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
