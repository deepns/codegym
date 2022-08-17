#include <asm-generic/errno.h>
#include <errno.h>
#include <fcntl.h>
#include <netdb.h>
#include <stdio.h>
#include <sys/poll.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <strings.h>
#include <unistd.h>
#include <string.h>
#include <openssl/ssl.h>
#include <openssl/err.h>
#include <poll.h>

int create_socket(char *server, int port)
{
    int sockfd = -1;
    struct hostent *host;
    struct sockaddr_in addr;
    if ((host = gethostbyname(server)) == NULL)
    {
        perror(server);
        abort();
    }

    printf("Connecting to %s:%d\n", server, port);

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = *(long*)(host->h_addr);

    sockfd = socket(PF_INET, SOCK_STREAM, 0);

    // get the file status flags, OR with additional
    // flags (e.g. O_NONBLOCK) and set the new flags
    int flags = fcntl(sockfd, F_GETFL, 0);
    fcntl(sockfd, F_SETFL, flags | O_NONBLOCK);

    if (connect(sockfd, (struct sockaddr*)&addr, sizeof(addr)) < 0) {
        if (errno == EWOULDBLOCK || errno == EINPROGRESS) {
            // TODO
            // What action to take for EWOULDBLOCK and EINPROGRESS errors?
            printf("Connecting...\n");
        } else {
            close(sockfd);
            perror("connect() failed");
        }
    }
    return sockfd;
}

SSL_CTX* create_context()
{
    // create SSL context
    // Using the general purpose client method.
    // The actual protocol version used will be negotiated to the highest
    // version mutually supported by the client and the server.
    // The supported protocols are SSLv3, TLSv1, TLSv1.1, TLSv1.2 and TLSv1.3.

    SSL_CTX *context = SSL_CTX_new(TLS_client_method());
    if (context == NULL) {
        ERR_print_errors_fp(stderr);
        return NULL;
    }

    if (SSL_CTX_load_verify_locations(context, "certs/root_ca.crt", NULL) != 1) {
        ERR_print_errors_fp(stderr);
        SSL_CTX_free(context);
        return NULL;
    }

    return context;
}

void show_help()
{
    printf("USAGE: ./tls_client_nb <server> <port>\n"
           "e.g. ./tls_client_nb 142.251.162.103 443\n");
}

void handle_ssl_connection(SSL *ssl)
{
    char buf[1024];
    bzero(buf, sizeof(buf));
    sprintf(buf, "%s", "Houston! We have a problem!");

    // TODO
    // Handle errors from SSL_write
    // If the write cannot be completed, check the error from SSL_get_error.
    // If SSL_ERROR_WANT_READ or SSL_ERROR_WANT_WRITE is set, wait accordingly.
    // May be if the buf is large enough, this will fail?
    size_t bytessent = 0;
    // can also use SSL_write() instead of SSL_write_ex()
    (void)SSL_write_ex(ssl, &buf, strlen(buf), &bytessent);
    printf("Sent msg:%s of bytessent:%ld\n", buf, bytessent);        
    sleep(1);
}

int do_ssl_connect(SSL *ssl, int sockfd)
{
    int rc;
    struct pollfd poll_fd = { .fd = sockfd };
    int poll_timeout_ms = 1000;
    while (1) {
        printf("Doing SSL_connect()...\n");
        rc = SSL_connect(ssl);
        if (rc == 1) {
           return 1;
        } else {
            int ssl_err = SSL_get_error(ssl, rc);
            if (ssl_err == SSL_ERROR_WANT_READ) {
                poll_fd.events = POLLIN;
            } else if (ssl_err == SSL_ERROR_WANT_WRITE) {
                poll_fd.events = POLLOUT;
            } else {
                fprintf(stderr, "ERROR: Hit unexpected error: rc=%d, ssl_err=%d\n", rc, ssl_err);
                return 0;
            }

            rc = poll(&poll_fd, 1, poll_timeout_ms);
            // TODO check for the right events in the descriptor
            if (rc == -1) {
                fprintf(stderr, "ERROR: Hit error in polling after SSL_connect failure. rc=%d, ssl_err=%d\n", rc, ssl_err);
                return 0;
            }
        }
    }
}

int main(int argc, char **argv)
{
    if (argc != 3) {
        fprintf(stderr, "Invalid arguments\n");
        show_help();
        exit(EXIT_FAILURE);
    }

    char *server = argv[1];
    int port = atoi(argv[2]); // can use strtol instead.
    if (port == 0) {
        fprintf(stderr, "Invalid port value: %s", argv[2]);
        exit(EXIT_FAILURE);
    }

    int sockfd = create_socket(server, port);
    if (sockfd < 0) {
        fprintf(stderr, "Failed to connect.. Quitting\n");
        exit(EXIT_FAILURE);
    }

    // Setup SSL stuff
    SSL_CTX *context = create_context();
    SSL *ssl = SSL_new(context);

    // link SSL object and the socket fd
    SSL_set_fd(ssl, sockfd);

    int connected = do_ssl_connect(ssl, sockfd);

    if (connected) {
        printf("SSL_connect succeeded. Connected to %s:%d\n", server, port);
        handle_ssl_connection(ssl);
    }

    SSL_CTX_free(context);
    SSL_free(ssl);
    close(sockfd);
    return 0;
}
