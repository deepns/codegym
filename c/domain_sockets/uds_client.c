#include <errno.h>
#include <netdb.h>
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <strings.h>
#include <unistd.h>
#include <string.h>
#include <sys/un.h>

#include "uds.h"

#define BUFFER_SIZE 256

void check_err(long err, char *errmsg);
void handle_connection(int sockfd); 

int main()
{
    int sockfd;
    struct sockaddr_un addr;
    
    bzero(&addr, sizeof(addr));
    addr.sun_family = AF_UNIX;

    strncpy(addr.sun_path, SOCKET_NAME, sizeof(addr.sun_path));

    sockfd = socket(PF_UNIX, SOCK_STREAM, 0);
    check_err(sockfd, "socket() failed");

    check_err(connect(sockfd, (struct sockaddr*)&addr, sizeof(addr)),
                      "connect() failed");

    handle_connection(sockfd);

    return EXIT_SUCCESS;
}

void handle_connection(int sockfd)
{
    char buffer[BUFFER_SIZE] = "\0";
    ssize_t byteswritten = -1;
    ssize_t bytesread = -1;

    snprintf(buffer, BUFFER_SIZE, "name: the_quick_brown_fox_jumps_over_lazy_dog");
    byteswritten = write(sockfd, buffer, strlen(buffer));
    check_err(byteswritten, "write() failed");
    printf("msg_to_server=%s, msglen=%ld\n", buffer, byteswritten);


    bytesread = read(sockfd, buffer, BUFFER_SIZE);
    check_err(bytesread, "read() failed");
    buffer[bytesread] = '\0';
    printf("msg_from_server=%s, msglen=%ld\n", buffer, bytesread);

    close(sockfd);
}

void check_err(long err, char *errmsg)
{
    if (err == -1) {
        perror(errmsg);
        exit(EXIT_FAILURE);
    }
}