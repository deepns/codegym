#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <sys/socket.h>
#include <unistd.h>

#define BUFFER_SIZE 128
#define SERVER_NAME "localhost"
#define SERVER_PORT 11011

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

static inline void check_err(int err, const char *msg)
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
    check_err(sockfd, "socket() failed");

    if (connect(sockfd, (const struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("connect() failed");
        close(sockfd);
        exit(EXIT_FAILURE);
    }

    printf("Connected to %s:%d\n", hostname, port);
    return sockfd;
}

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

    // read/write calls return -1 on error.
    // ignoring them for now.

    for (int i=0; i<4; i++) {
        byteswritten = write(sockfd, (const void *)msgs[i], strlen(msgs[i]));
        bytesread = read(sockfd, (void *)&readbuf, BUFFER_SIZE);
        readbuf[bytesread] = '\0';
        printf("wrote:%s, read:%s\n", msgs[i], readbuf);
        sleep(1);
    }

    snprintf(writebuf, BUFFER_SIZE, "%s", "bye");
    byteswritten = write(sockfd, (const void *)writebuf, 3);
    close(sockfd);
}

int main(int argc, char **argv)
{
    int serversock = create_connection(SERVER_NAME, SERVER_PORT);
    process_connection(serversock);
    close(serversock);
    return 0;
}