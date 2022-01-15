#include <stdio.h>
#include <netdb.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <strings.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <errno.h>
#include <sys/fcntl.h>

/**
 * @brief Create a connection object
 *
 * @param hostname
 * @param port
 * @return int socket descriptor to the server connection
 */
int create_connection(const char *hostname, int port);

/**
 * @brief Process connection to the server
 *
 * @param sockfd socket descriptor to the server
 */
void process_connection(int sockfd);

/**
 * @brief Exit with EXIT_FAILURE if err indicates fatal ( < 0)
 * 
 * @param err 
 * @param msg 
 */
static inline void check_err_exit(int err, const char *msg)
{
    if (err < 0) {
        perror(msg);
        exit(EXIT_FAILURE);
    }
}

int create_connection(const char *hostname, int port)
{
    struct hostent *host;
    struct sockaddr_in addr;
    if ((host = gethostbyname(hostname)) == NULL) {
        herror("gethostbyname() failed");
        exit(EXIT_FAILURE);
    }

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = *(long *)(host->h_addr);
    addr.sin_port = htons(port);

    int sockfd = socket(PF_INET, SOCK_STREAM, 0);
    check_err_exit(sockfd, "socket() failed");

    if (connect(sockfd, (const struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("connect() failed");
        close(sockfd);
        exit(EXIT_FAILURE);
    }

    printf("Connected to %s:%d\n", hostname, port);
    return sockfd;
}

#define BUFFER_SIZE 128

void process_connection(int sockfd)
{
    char writebuf[BUFFER_SIZE] = "\0";
    char readbuf[BUFFER_SIZE] = "\0";

    ssize_t byteswritten;
    ssize_t bytesread;

    char *msgs[4] = {
        "hello",
        "some message",
        "another message",
        "one more"
    };

    for (int i=0; i<4; i++) {
        // TODO
        // writing very small buffers here, so not worrying
        // about partial writes though socket is set in non-blocking mode.
        // when sending large chunks of data, write() may return
        // early with data still left to write.
        byteswritten = write(sockfd, (const void *)msgs[i], strlen(msgs[i]));
        while (1) {
            bytesread = read(sockfd, (void *)&readbuf, BUFFER_SIZE);
            if (bytesread == -1) {
                if (errno == EAGAIN) {
                    printf("waiting for response from server\n");
                    sleep(2);
                }
                continue;
            } else if (bytesread == 0) {
                printf("server closed the connection\n");
                break;
            } else {
                readbuf[bytesread] = '\0';
                printf("wrote:%s, read:%s\n", msgs[i], readbuf);
                break;
            }
            sleep(1);
        }
    }

    snprintf(writebuf, BUFFER_SIZE, "%s", "bye");
    byteswritten = write(sockfd, (const void *)writebuf, 3);
    close(sockfd);
}

#define SERVER_NAME "localhost"
#define SERVER_PORT 9699

int main(int argc, char **argv)
{
    int serverconn = create_connection(SERVER_NAME, SERVER_PORT);

    // TODO Why does a client want to connect in non-blocking mode?
    // what are some practical use cases for that?
    // Turning on non-blocking mode
    int flags = fcntl(serverconn, F_GETFL, 0);
    fcntl(serverconn, F_SETFL, flags | O_NONBLOCK);

    process_connection(serverconn);
    close(serverconn);
    return 0;
}