#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <sys/socket.h>
#include <sys/select.h>
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
void handle_connection(int sockfd, int *closed);

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

void handle_connection(int sockfd, int *closed)
{
    char buf[BUFFER_SIZE] = "\0";
    ssize_t bytesread = read(sockfd, (void *)buf, BUFFER_SIZE);

    buf[bytesread] = '\0';
    if (strcmp(buf, "bye") == 0) {
        printf("Closing the connection...bye\n");
        close(sockfd);
        if (closed) {
            *closed = 1;
        }
    } else {
        printf("client sent: %s\n", buf);
        write(sockfd, (const void *)buf, bytesread);
    }
}

int main(int argc, char **argv)
{
    int serversock = setup_server(SERVER_PORT, SERVER_BACKLOG);
    int fd_max = serversock;
    fd_set listening_sockets;
    fd_set active_sockets;
    int err = -1;

    FD_ZERO(&listening_sockets);
    FD_ZERO(&active_sockets);
    FD_SET(serversock, &active_sockets);

    for (;;) {
        // On success, select returns the number of ready descriptors that are contained in the descriptor sets
        // On failure, select returns -1
        // On time limit expiration, select returns 0.
        // NOTE: select resets the fd sets. so reset the listening sockets
        // with the active sockets
        listening_sockets = active_sockets;

        // The descriptors 0 through fd_max - 1 are examined during select.
        if ((err = select(fd_max + 1, &listening_sockets, NULL, NULL, NULL)) < 0) {
            check_err(err, "select() failed");
        }
            
        // NOTE: checking each socket until FD_SETSIZE can be expensive
        // Hence restricting to the max descriptor seen so far.
        // FD_SETSIZE can be vary based on the platform.
        // typically 1024 per process.
        for (int fd=0; fd <= fd_max; fd++) {
            if (FD_ISSET(fd, &listening_sockets)) {
                if (fd == serversock) {
                    int clientsock = accept_connection(serversock);
                    // add new client to set of active connections
                    FD_SET(clientsock, &active_sockets);
                    // update max fd seen so far.
                    if (clientsock > fd_max) {
                        fd_max = clientsock;
                    }
                } else {
                    int is_closed = 0;
                    handle_connection(fd, &is_closed);
                    // if the connection was closed in the last handling,
                    // clear it from the list of active sockets.
                    if (is_closed) {
                        FD_CLR(fd, &active_sockets);
                    }
                }
            }
        }
    }

    // NOT_REACHED
    close(serversock);
    return 0;
}