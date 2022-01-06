#include <errno.h>
#include <fcntl.h>
#include <netdb.h>
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <strings.h>
#include <unistd.h>
#include <string.h>

int main()
{
    char *SERVER = "127.0.0.1";
    int port = 9699;

    int sockfd;
    struct hostent *host;
    struct sockaddr_in addr;
    if ((host = gethostbyname(SERVER)) == NULL)
    {
        perror(SERVER);
        abort();
    }

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = *(long*)(host->h_addr);

    sockfd = socket(PF_INET, SOCK_STREAM, 0);

    // get the file status flags, OR with additional
    // flags (e.g. O_NONBLOCK) and set the new flags
    int flags = fcntl(sockfd, F_GETFL, 0);
    fcntl(sockfd, F_SETFL, flags | O_NONBLOCK);

    int err = -1;
    if ((err = connect(sockfd, (struct sockaddr*)&addr, sizeof(addr))) < 0) {
        printf("Connected to socket----, err = %d\n", err);   
        if (err != EWOULDBLOCK || err != EINPROGRESS) {
            close(sockfd);
            perror("connect() failed");
            exit(EXIT_FAILURE);
        }
        printf("Connected to socket++++++++++   ??\n");   
    }

	char buf[1024];
	bzero(buf, sizeof(buf));
    sprintf(buf, "%s", "client: foo");

    ssize_t bytessent = send(sockfd, &buf, strlen(buf), 0);
    printf("Sent msg:%s of bytessent:%ld\n", buf, bytessent);

	bzero(buf, sizeof(buf));
	ssize_t bytesread = recv(sockfd, &buf, sizeof(buf), 0);
    printf("Received: Message (%s) of length %ld from %s\n", buf, bytesread, SERVER);

    close(sockfd);
    return 0;
}
