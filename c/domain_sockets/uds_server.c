#include <netdb.h> // for sockaddr, INADDR_ANY
#include <stdio.h>
#include <sys/socket.h> // for socket()
#include <stdlib.h> // for EXIT_FAILURE
#include <strings.h> // for bzero
#include <unistd.h> // for close()
#include <string.h> // for strlen
#include <sys/un.h> // for sockaddr_un

#include "uds.h"

#define BUFFER_SIZE 256

/**
 * @brief Set the up server's domain socket
 * 
 * @param socket_name path to socket
 * @param backlog number of backlog connections
 * @return socket descriptor 
 */
int setup_socket(const char *socket_name, int backlog);

/**
 * @brief Handle client connection
 * 
 * @param client_id client identifier
 * @param sockfd socket descriptor
 */
void handle_client(int client_id, int sockfd);

/**
 * @brief Check error code and exit if error is fatal (-1)
 * 
 * @param err 
 * @param errmsg 
 */
void check_err(int err, char *errmsg);

int main()
{
    int serversock = setup_socket(SOCKET_NAME, 2/* backlog */);
    int client_id = 0;

    // listen for connections and accept
    while (1) {
        /*
         * for network sockets, we usually specify the addr and length
         * to get the client address details. Since domain sockets are
         * used only within the same host, addr is not required.
         * If we do pass one, accept() sets the address family (AF_UNIX) only.
         * (gdb) p caddr
            $1 = {
                sun_family = 1,
                sun_path = '\000' <repeats 107 times>
            }
         */
        int clientsockfd = accept(serversock, NULL /* addr */, NULL /* addr len */);
        check_err(clientsockfd, "accept() failed");

        handle_client(++client_id, clientsockfd);
    }

    close(serversock);
    unlink(SOCKET_NAME);
    return 0;
}

int setup_socket(const char *socket_name, int backlog)
{
    int sockfd;

    // Use sockaddr_un for UNIX domain socket addresses
    struct sockaddr_un addr;
    bzero(&addr, sizeof(addr));
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, socket_name, sizeof(addr.sun_path));

    // clear the socket file if exists
    unlink(socket_name);

    // create the socket
    sockfd = socket(PF_UNIX, SOCK_STREAM, 0);
    check_err(sockfd, "Failed to create socket");
    
    // bind
    check_err(bind(sockfd, (struct sockaddr *)&addr, sizeof(addr)),
              "bind() failed");

    // listen
    check_err(listen(sockfd, backlog),
              "Failed to enable listening mode on the socket");
    
    return sockfd;
}

void handle_client(int client_id, int sockfd)
{
    char buffer[BUFFER_SIZE] = "\0";
    ssize_t bytesread = -1;
    ssize_t byteswritten = -1;

    printf("client_id=%d\n", client_id);

    bytesread = read(sockfd, buffer, BUFFER_SIZE);
    check_err(bytesread, "read() failed");
    buffer[bytesread] = '\0';
    printf("msg_from_client=%s, msglen=%ld\n", buffer, bytesread);

    snprintf(buffer, BUFFER_SIZE, "Hello, there. You are client#%d", client_id);
    byteswritten = write(sockfd, buffer, strlen(buffer));
    check_err(byteswritten, "write() failed");
    printf("msg_to_client=%s, msglen=%ld\n", buffer, byteswritten);

    close(sockfd);
}

void check_err(int err, char *errmsg)
{
    if (err == -1) {
        perror(errmsg);
        exit(EXIT_FAILURE);
    }
}