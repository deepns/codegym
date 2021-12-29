#include <netdb.h> // for sockaddr, INADDR_ANY
#include <stdio.h>
#include <sys/socket.h> // for socket()
#include <stdlib.h> // for EXIT_FAILURE
#include <strings.h> // for bzero
#include <unistd.h> // for close()
#include <string.h> // for strlen
#include <sys/un.h> // for sockaddr_un

#include "uds.h"

int create_server_uds_socket()
{
    int sockfd;

    // Use sockaddr_un for UNIX domain socket addresses
    struct sockaddr_un addr;
    bzero(&addr, sizeof(addr));
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, SOCKET_NAME, sizeof(addr.sun_path));

    // clear the existing socket file
    unlink(SOCKET_NAME);

    // create the socket
    sockfd = socket(PF_UNIX, SOCK_STREAM, 0);
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
    int serversock = create_server_uds_socket();
    int client_id = 0;

    // listen for connections and accept
    while (1) {
        int clientsockfd = accept(serversock, NULL, NULL);

        if (clientsockfd < 0) {
            perror("failed to accept client");
            close(serversock);
            exit(EXIT_FAILURE);
        }

        client_id++;

        char clientdata[1024] = "\0";
        int bytesrecv = recv(clientsockfd, &clientdata, sizeof(clientdata), 0 /* flags */);
        printf("client_id=%d, msg=%s, msg_length=%d\n", client_id, clientdata, bytesrecv);

        char messagetoclient[2048];
        sprintf(messagetoclient, "Hello, there! You are client#%d", client_id);
        int bytessent = send(clientsockfd, &messagetoclient, strlen(messagetoclient), 0 /* flags */);

        close(clientsockfd);
    }

    close(serversock);
    return 0;
}