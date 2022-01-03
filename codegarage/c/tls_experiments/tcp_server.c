#include <netdb.h> // for sockaddr, INADDR_ANY
#include <stdio.h>
#include <sys/socket.h> // for socket()
#include <stdlib.h> // for EXIT_FAILURE
#include <strings.h> // for bzero
#include <unistd.h> // for close()
#include <string.h>

int create_server_socket(int port)
{
    int sockfd;
    struct sockaddr_in addr;

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    // create the socket
    sockfd = socket(PF_INET, SOCK_STREAM, 0);
    if (sockfd < 0) {
        perror("Failed to create socket");
        exit(EXIT_FAILURE);
    }
    
    // bind
    if (bind(sockfd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("Failed to bind the socket");
        exit(EXIT_FAILURE);
    }

    // listen
    if (listen(sockfd, 2 /*backlog*/) < 0) {
        perror("Failed to enable listening mode on the socket");
        exit(EXIT_FAILURE);
    }
    
    return sockfd;
}

int main()
{
    int port = 9699;
    int serversock = create_server_socket(port);

    // listen for connections and accept
    while (1) {
        struct sockaddr_in addr;
        unsigned int len = sizeof(addr);
        
        int clientsockfd = accept(serversock, (struct sockaddr *)&addr, &len);
        if (clientsockfd < 0) {
            perror("failed to accept client");
            close(serversock);
            exit(EXIT_FAILURE);
        }

        printf("client_port=%d\n", addr.sin_port);

        char clientdata[1024] = "\0";
        int bytesrecv = recv(clientsockfd, &clientdata, sizeof(clientdata), 0 /* flags */);
        printf("Received msg:%s of length:%d from client\n", clientdata, bytesrecv);

        char messagetoclient[2048];
        sprintf(messagetoclient, "hello, there! You sent (%s)", clientdata);
        int bytessent = send(clientsockfd, &messagetoclient, strlen(messagetoclient), 0 /* flags */);

        close(clientsockfd);
    }

    close(serversock);
    return 0;
}