#include <errno.h>
#include <netdb.h>
#include <openssl/ssl.h>
#include <openssl/err.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#include "cert_util.h"

SSL_CTX* create_context_mtls(const char *certfile,
                            const char *keyfile,
                            const char *cafile,
                            const char *capath)
{
    // create SSL context
    // Using the general purpose client method.
    // The actual protocol version used will be negotiated to the highest
    // version mutually supported by the client and the server.
    // The supported protocols are SSLv3, TLSv1, TLSv1.1, TLSv1.2 and TLSv1.3.

    // There is a slight difference in behavior between TLSv1.2 and TLSv1.3
    // in terms of validating the certificate during handshake.
    // In TLSv1.3, SSL_connect() would succeed even if the server rejects
    // client certificate (for e.g. client's cert expired, or invalid, or didn't
    // present at all/). First read from the client application would then
    // fail since the underlying connection is already closed by the server.
    // In case of TLSv1.2, SSL_connect() would fail if there is a failure
    // in TLS handshake. 
    // This [issue](https://github.com/openssl/openssl/issues/8500)
    // has some good comments explaining this version difference.

    // In OpenSSL 1.1.1, usage of individual TLS method is deprecated
    // using a specific version (for e.g. TLSv1_2_client_method()) will
    // raise the deprecated warning during compilation.
    const SSL_METHOD *method = TLS_client_method();
    SSL_CTX *context = SSL_CTX_new(method);
    if (context == NULL) {
        ERR_print_errors_fp(stderr);
        return NULL;
    }

    if (context) {
        /*
        * All the SSL_CTX_xxx functions below return 1 on success.
        */
        if ((SSL_CTX_use_certificate_chain_file(context, certfile) == 1) &&
            (SSL_CTX_use_PrivateKey_file(context, keyfile, SSL_FILETYPE_PEM) == 1) &&
            (SSL_CTX_load_verify_locations(context, cafile, capath) == 1)) {

            int verify_flags = SSL_VERIFY_PEER |
                               SSL_VERIFY_FAIL_IF_NO_PEER_CERT;
            SSL_CTX_set_verify(context, verify_flags, NULL /* verify callback */);
            return context;
        }
    }

    ERR_print_errors_fp(stderr);
    if (context) {
        SSL_CTX_free(context);
    }
    return NULL;
}

int create_socket(const char *server, int port)
{
    int sockfd = -1;
    struct hostent *host;
    struct sockaddr_in addr;
    if ((host = gethostbyname(server)) == NULL) {
        perror("gethostbyname failed");
        return -1;
    }

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = *(long*)(host->h_addr);

    sockfd = socket(PF_INET, SOCK_STREAM, 0);
    if (sockfd < 0) {
        perror("Failed to create socket");
        return -1;
    }

    if (connect(sockfd, (struct sockaddr*)&addr, sizeof(addr)) != 0) {
        perror("Failed to connect to server");
        close(sockfd);
        return -1;
    }

    return sockfd;
}

int main()
{
    /*
     * Initialize SSL library
     * linking with only openssl version 1.1.0 above
     * so no explicit initialization and free of the library
     * is required. However, applications can still call the init
     * explicitly depending on the needs.
     */
    // (void)OPENSSL_init_ssl(0, NULL);

    char *SERVER = "localhost";
    int port = 9899;
    int sockfd = -1;
    const char * CLIENT_CERT = "certs/client.crt";
    const char * CLIENT_KEY = "certs/client.key";
    const char * CAFILE = "certs/root_ca.crt";

    sockfd = create_socket(SERVER, port);
    if (sockfd == -1) {
        exit(EXIT_FAILURE);
    }

    SSL_CTX *context = create_context_mtls(
                                    CLIENT_CERT, 
                                    CLIENT_KEY,
                                    CAFILE,
                                    NULL);
    if (context == NULL) {
        exit(EXIT_FAILURE);
    }

    // Wrap the socket with SSL
    SSL *ssl = SSL_new(context);

    // attach the ssl object to the socket
    SSL_set_fd(ssl, sockfd);

    /* 
     * securely connect with SSL
     * The behavior of TLS handshake and acceptance is different between
     * TLS versions 1.2 and 1.3. SSL_connect() would succeed even if server
     * fails to accept the connection due to a failure in verifying the
     * client certificate. See my comments in create_context()
     */
    if (SSL_connect(ssl) != 1)
    {
        fprintf(stderr, "Failed to connect to server\n");
        ERR_print_errors_fp(stderr);
        SSL_free(ssl);
        SSL_CTX_free(context);
        exit(EXIT_FAILURE);
    }

    printf("Showing server certificate:\n");
    show_cert_info(SSL_get_peer_certificate(ssl));

	char buf[1024];
    bzero(buf, sizeof(buf));
    sprintf(buf, "%s", "client: foo");

    /*
     * typically client writes to the server first before reading
     * anything from the server. In case of TLS1.3, if there is a
     * certificate verification failure, SSL_connect() would still
     * succeed and subsequently SSL_write() would also succeed even
     * though server had not accepted the connection fully. The buffer
     * will be written to the socket, but server would not read it.
     * However, SSL_read would fail if the server hadn't accepted
     * the connection.
     */

    size_t bytessent = 0;
    int write_err = SSL_write_ex(ssl, &buf, strlen(buf), &bytessent);
    if (write_err <= 0) {
        fprintf(stderr, "failed to write: err:%d, ssl_err=%d\n", write_err, SSL_get_error(ssl, write_err));
    } else {
        printf("Sent msg:%s of bytessent:%ld\n", buf, bytessent);
    }

	bzero(buf, sizeof(buf));
	size_t bytesread = 0;
    int ret = SSL_read_ex(ssl, &buf, sizeof(buf), &bytesread);
    if (ret <= 0) {
        fprintf(stderr, "Received read_error:%d, ssl_error_code=%d\n", ret, SSL_get_error(ssl, ret));
    } else {
        printf("Received %s of length %ld from %s\n", buf, bytesread, SERVER);
    }

    SSL_free(ssl);
    SSL_CTX_free(context);
    close(sockfd);
    return 0;
}
