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

int main()
{
    int sockfd;
    struct sockaddr_un addr;
    bzero(&addr, sizeof(addr));
    addr.sun_family = AF_UNIX;

    strncpy(addr.sun_path, SOCKET_NAME, sizeof(addr.sun_path));
    
    sockfd = socket(PF_UNIX, SOCK_STREAM, 0);
    if (connect(sockfd, (struct sockaddr*)&addr, sizeof(addr)) != 0) {
        close(sockfd);
        exit(EXIT_FAILURE);
    }

	char buf[1024];
	bzero(buf, sizeof(buf));
    sprintf(buf, "%s", "client: foo");

    ssize_t bytessent = send(sockfd, &buf, strlen(buf), 0);
    printf("msg_sent=%s, bytessent=%ld\n", buf, bytessent);

	bzero(buf, sizeof(buf));
	ssize_t bytesread = recv(sockfd, &buf, sizeof(buf), 0);
    printf("message_from_server=%s, msg_length=%ld\n", buf, bytesread);

    close(sockfd);
    return 0;
}
