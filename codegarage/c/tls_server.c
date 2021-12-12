#include <stdlib.h>
#include <stdio.h>
#include <sys/socket.h>
#include <netdb.h>
#include <unistd.h>
#include <openssl/ssl.h>
#include <openssl/err.h>
#include <string.h>

int create_server_socket(int port)
{
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    int sockfd = socket(
                    PF_INET /* IPv4 communication domain. */, 
                    SOCK_STREAM, /* socket type - TCP */
                    0 /* protocol to use. 0=IP. see /etc/protocols */
                    );
    
    if (sockfd < 0) {
        perror("Unable to create socket");
        exit(EXIT_FAILURE);
    }

    if (bind(sockfd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("Unable to bind socket");
        exit(EXIT_FAILURE);
    }

    if (listen(sockfd, 2) < 0) {
        perror("Unable to listen");
        exit(EXIT_FAILURE);
    }

    return sockfd;
}

SSL_CTX *create_ssl_context(const char *certfile, const char *keyfile)
{
    SSL_CTX *context = SSL_CTX_new(TLS_server_method());
    if (context == NULL) {
        ERR_print_errors_fp(stderr);
        return NULL;
    }

    /* load the certificate and private keys into the context */
    if (SSL_CTX_use_certificate_chain_file(context, certfile) != 1) {
        ERR_print_errors_fp(stderr);
        SSL_CTX_free(context);
        return NULL;
    }

    if (SSL_CTX_use_PrivateKey_file(context, keyfile, SSL_FILETYPE_PEM) != 1) {
        ERR_print_errors_fp(stderr);
        SSL_CTX_free(context);
        return NULL;
    }

    return context;
}

int main(void)
{
    int sockfd;
    int port = 9090;
    SSL_CTX *context = NULL;
    const char *certfile = "certs/server.crt";
    const char *keyfile = "certs/server.key";

    // Initialize SSL library
    // linking with only openssl version 1.1.0 above
    // so no explicit initialization and free of the library
    // is required.
    // (void)OPENSSL_init_ssl(0, NULL);

    sockfd = create_server_socket(port);
    context = create_ssl_context(certfile, keyfile);
    if (context == NULL) {
        perror("failed to create SSL context");
        close(sockfd);
    }

    int clientId = 0;
    // Start listening
    while (1) {
        struct sockaddr_in addr;
        unsigned int len = sizeof(addr);

        int clientsockfd = accept(sockfd, (struct sockaddr *)&addr, &len);

        if (clientsockfd < 0) {
            perror("Unable to accept client connection");
            close(sockfd);
            exit(EXIT_FAILURE);
        }

        // create a ssl object and associate it with the client socket descriptor
        SSL *ssl = SSL_new(context);
        SSL_set_fd(ssl, clientsockfd);

        // do a SSL accept to accept the client connect.
        if (SSL_accept(ssl) <= 0) {
            /*
             * 0 -> The TLS/SSL handshake was not successful but was shut down
             * controlled and by the specifications of the TLS/SSL protocol
             * 
             * < 0 -> The TLS/SSL handshake was not successful because a fatal
             * error occurred either at the protocol level or a connection 
             * failure occurred. The shutdown was not clean.
             */
            ERR_print_errors_fp(stderr);
        } else {
            clientId++;
            printf("Connected to client#%d\n", clientId);

            // process the client connection
            char clientdata[1024] = "\0";
            size_t bytesread = 0;

            // TODO get the return value from SSL_read and SSL_write
            // and call SSL_get_error() to find more about the error.

            // read through the SSL object
            if (SSL_read_ex(ssl, &clientdata, sizeof(clientdata), &bytesread) <= 0) {
                perror("Failed to read from client");
            }

            printf("Received: msg:(%s) of length (%lu)\n", clientdata, bytesread);

            // send some msg to client over SSL
            char msg[1024] = "\0";
            sprintf(msg, "You are client#%d", clientId);
            if (SSL_write(ssl, &msg, strlen(msg)) <= 0) {
                perror("Failed to write to client");
            }
        }
        SSL_free(ssl);
        close(clientsockfd);
    }

    close(sockfd);
    SSL_CTX_free(context);
    return 0;
}