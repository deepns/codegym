#include <errno.h>
#include <netdb.h>
#include <openssl/ssl.h>
#include <openssl/err.h>
#include <stdio.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

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

int create_socket(const char *server, int port)
{
    int sockfd;
    struct hostent *host;
    struct sockaddr_in addr;
    if ((host = gethostbyname(server)) == NULL) {
        perror(server);
        exit(EXIT_FAILURE);
    }

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = *(long*)(host->h_addr);

    sockfd = socket(PF_INET, SOCK_STREAM, 0);
    if (connect(sockfd, (struct sockaddr*)&addr, sizeof(addr)) != 0) {
        perror(server);
        close(sockfd);
        exit(EXIT_FAILURE);
    }

    return sockfd;
}

/*
 * Print some info about the cert in the given SSL object
 */
void show_certs(SSL *ssl)
{
    X509 *cert;

    cert = SSL_get_peer_certificate(ssl);
    if (cert) {
        X509_NAME *issuer = X509_get_issuer_name(cert);
        X509_NAME *subject = X509_get_subject_name(cert);
        char *info;
        info = X509_NAME_oneline(issuer, 0, 0);
        printf("Cert issued by: %s\n", info);
        free(info);

        info = X509_NAME_oneline(subject, 0, 0);
        printf("Cert subject is: %s\n", info);
        free(info);
        X509_free(cert);
    } else {
        printf("Unable to get peer certificate\n");
    }
}

int main()
{
    // Initialize SSL library
    // linking with only openssl version 1.1.0 above
    // so no explicit initialization and free of the library
    // is required.
    // (void)OPENSSL_init_ssl(0, NULL);

    char *SERVER = "localhost";
    int port = 9090;
    int sockfd = create_socket(SERVER, port);

    SSL_CTX *context = create_context();
    if (context == NULL) {
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

    show_certs(ssl);

	char buf[1024];
	bzero(buf, sizeof(buf));
    sprintf(buf, "%s", "client: foo");

    size_t bytessent = 0;
    // can also use SSL_write() instead of SSL_write_ex()
    (void)SSL_write_ex(ssl, &buf, strlen(buf), &bytessent);
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
