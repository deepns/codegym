#include <errno.h>
#include <netdb.h>
#include <openssl/ssl.h>
#include <openssl/err.h>
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <strings.h>
#include <unistd.h>

int main()
{
    // Initialize SSL library
    (void)OPENSSL_init_ssl(0, NULL);

    // create SSL context
    // Using the general purpose client method.
    // The actual protocol version used will be negotiated to the highest
    // version mutually supported by the client and the server.
    // The supported protocols are SSLv3, TLSv1, TLSv1.1, TLSv1.2 and TLSv1.3.
    SSL_CTX *context = SSL_CTX_new(TLS_client_method());
    if (context == NULL)
    {
        ERR_print_errors_fp(stderr);
        exit(EXIT_FAILURE);
    }

    char *SERVER = "localhost";
    int port = 9090;

    int sockfd;
    struct hostent *host;
    struct sockaddr_in addr;
    if ((host = gethostbyname(SERVER)) == NULL)
    {
        perror(SERVER);
        exit(EXIT_FAILURE);
    }

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = *(long*)(host->h_addr);

    sockfd = socket(PF_INET, SOCK_STREAM, 0);
    if (connect(sockfd, (struct sockaddr*)&addr, sizeof(addr)) != 0)
    {
        close(sockfd);
        perror(SERVER);
        exit(EXIT_FAILURE);
    }

    // Wrap the socket with SSL
    SSL *ssl = SSL_new(context);

    // attach the ssl object to the socket
    SSL_set_fd(ssl, sockfd);

    // securely connect with SSL
    if (SSL_connect(ssl) != 1)
    {
        ERR_print_errors_fp(stderr);
        SSL_free(ssl);
        SSL_CTX_free(context);
        exit(EXIT_FAILURE);
    }

	char buf[1024];
	bzero(buf, sizeof(buf));
    sprintf(buf, "%s", "client: foo");

    size_t bytessent = 0;
    // can also use SSL_write() instead of SSL_write_ex()
    (void)SSL_write_ex(ssl, &buf, sizeof(buf), &bytessent);
    printf("Sent msg:%s of bytessent:%ld\n", buf, bytessent);

	bzero(buf, sizeof(buf));
	size_t bytesread = 0;
    // SSL_read_ex provides the actual number of bytes read
    SSL_read_ex(ssl, &buf, sizeof(buf), &bytesread);
    printf("Received %s of length %ld from %s\n", buf, bytesread, SERVER);

    SSL_free(ssl);
    close(sockfd);
    SSL_CTX_free(context);
    return 0;
}
