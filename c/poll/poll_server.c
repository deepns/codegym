#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <sys/socket.h>
#include <poll.h>
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

#define MAX_CLIENTS 2

int main(int argc, char **argv)
{
    int serversock = setup_server(SERVER_PORT, SERVER_BACKLOG);
    struct pollfd fds[MAX_CLIENTS];
    int nclients = 0;

    bzero(&fds, sizeof(struct pollfd) * MAX_CLIENTS);

    while (nclients < MAX_CLIENTS) {
        int clientsock = accept_connection(serversock);
        fds[nclients].fd = clientsock;
        fds[nclients].events = POLLIN;
        nclients++;
    }

    int n_ready_fds = 0;
    while (1) {
        // On success, poll returns the number of fds ready for I/O
        n_ready_fds = poll(fds,
                           MAX_CLIENTS /* number of fds to poll */, 
                           30000 /* timeout in millisecs */);

        if (n_ready_fds == 0) {
            printf("poll timer expired...polling again\n");
            continue;
        } else if (n_ready_fds == -1) {
            check_err(n_ready_fds, "poll() failed");
        } else {
            // Check which FDs are ready for reading
            for (int i=0; i<MAX_CLIENTS; i++) {
                if (fds[i].revents & POLLIN) {
                    int is_closed = 0;
                    handle_connection(fds[i].fd, &is_closed);
                    if (is_closed) {
                        nclients--;
                    }
                }
            }
            // if all client connections are closed, close
            // the server and exit.
            if (!nclients) {
                break;
            }
        }
    }

    // NOT_REACHED
    close(serversock);
    return 0;
}