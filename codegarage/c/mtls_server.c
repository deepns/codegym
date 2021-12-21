#include <openssl/ssl.h>
#include <openssl/err.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

/*
 * Create a IPv4 socket at the given port.
 * currently listening on any addr on the host.
 */
int create_server_socket(int port)
{
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    int sockfd = socket(
        PF_INET /*protocol family*/,
        SOCK_STREAM /*socket type*/,
        0 /*protocol to use. 0=IP*/
        );
    
    if (sockfd < 0) {
        perror("Unable to create a socket");
        exit(EXIT_FAILURE);
    }

    if (bind(sockfd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("Unable to bind the socket");
        exit(EXIT_FAILURE);
    }

    if (listen(sockfd, 2) < 0) {
        perror("Unable to listen");
        exit(EXIT_FAILURE);
    }

    return sockfd;
}

/*
 * Create the SSL context and initialize it with the
 * parameters required to do mutual TLS authentication.
 */
SSL_CTX *create_context_mtls(
    const char *certfile,
    const char *keyfile,
    const char *cafile,
    const char *capath)
{
    /*
     * note: SSL_CTX object is reference counted.
     * Using the general purpose server method. The actual TLS version
     * will be selected during negotiation.
     * e.g. of specific versions TLSv1_2_server_method(), TLSv1_1_server_method(), 
     */
    SSL_CTX *ctx = SSL_CTX_new(TLS_server_method());
    if (ctx) {
        /*
        * All the SSL_CTX_xxx functions below return 1 on success.
        */
        if ((SSL_CTX_use_certificate_chain_file(ctx, certfile) == 1) &&
            (SSL_CTX_use_PrivateKey_file(ctx, keyfile, SSL_FILETYPE_PEM) == 1) &&
            (SSL_CTX_load_verify_locations(ctx, cafile, capath) == 1)) {

            int verify_flags = SSL_VERIFY_PEER |
                               SSL_VERIFY_FAIL_IF_NO_PEER_CERT |
                               SSL_VERIFY_CLIENT_ONCE;
            SSL_CTX_set_verify(ctx, verify_flags, NULL /* verify callback */);
            return ctx;
        }
        // hit some error while updating the context with certs
        SSL_CTX_free(ctx);
    }
    
    ERR_print_errors_fp(stderr);
    return NULL;
}

void recv_from_client(SSL *ssl, int sockfd, char *buf, int buflen)
{
    if (ssl) {
        size_t bytesread;
        // SSL_read doesn't return the number of bytes read. so using SSL_read_ex
        int readerr = SSL_read_ex(ssl, buf, buflen, &bytesread);
        if (readerr <= 0) {
            fprintf(stderr, 
                    "Hit error=%d from SSL_read, ssl_error=%d\n",
                    readerr, SSL_get_error(ssl, readerr));
        } else {
            printf("Received msg=%s of length=%lu from client over SSL\n", buf, bytesread);
        }
    } else {
        ssize_t bytesreceived = recv(sockfd, buf, buflen, 0 /*flags*/);
        printf("Received msg=%s of length=%ld from client\n", buf, bytesreceived);
    }
}

void send_to_client(SSL *ssl, int sockfd, char *buf, int buflen)
{
   if (ssl) {
        size_t byteswritten;
        int writeerr = SSL_write_ex(ssl, buf, buflen, &byteswritten);
        if (writeerr <= 0) {
            fprintf(stderr,
                    "Hit error=%d from SSL_write, ssl_error=%d\n",
                    writeerr, SSL_get_error(ssl, writeerr));
        } else {
            printf("Written %lu bytes of msg=%s over SSL\n", byteswritten, buf);
        }
    } else {
        ssize_t bytessent = send(sockfd, buf, buflen, 0);
        printf("Sent %ld bytes from msg=%s\n", bytessent, buf);
    }
}

void show_cert_info(SSL *ssl)
{
    X509 *cert = NULL;
    
    cert = SSL_get_peer_certificate(ssl);
    if (!cert) {
        fprintf(stderr, "Unable to get cert from SSL object\n");
    }

    /* 
     * X509_NAME_oneline produces non-standard output and can be
     * inconsistent at times. So its usage is discouraged in new
     * applications. Use X509_NAME_print_ex/X509_NAME_print_ex_fp
     * instead.
     * https://www.openssl.org/docs/man1.1.1/man3/X509_NAME_oneline.html
     */
    
    // X509_NAME *issuer = X509_get_issuer_name(cert);
    // char *issuer_info = X509_NAME_oneline(issuer, NULL, 0);

    // if (issuer_info) {
    //     printf("Cert issued by: %s\n", issuer_info);
    //     free(issuer_info);
    // }

    // X509_NAME *subject = X509_get_subject_name(cert);
    // char *subject_info = X509_NAME_oneline(subject, NULL, 0);
    // if (subject_info) {
    //     printf("Cert subject: %s\n", subject_info);
    // }

    X509_NAME *issuer = X509_get_issuer_name(cert);
    X509_NAME *subject = X509_get_subject_name(cert);

    printf("Client Cert issued by:\n");
    X509_NAME_print_ex_fp(stdout, issuer, 2 /*indent*/, XN_FLAG_ONELINE);
    printf("\n");

    printf("Client Cert subject:\n");
    X509_NAME_print_ex_fp(stdout, subject, 2 /*indent*/, XN_FLAG_MULTILINE);
    printf("\n");

    // refcnt taken on pubkey. so must be freed up after use
    // EVP_PKEY *pubkey = X509_get_pubkey(cert);
    //EVP_PKEY_free(pubkey);

    X509_free(cert);
}

int main()
{
    int port = 9899;
    int serversock = create_server_socket(port);
    int clientId = 0;

    const char *certfile = "certs/server.crt";
    const char *keyfile = "certs/server.key";
    const char *cafile = "certs/root_ca.crt";

    SSL_CTX *context = create_context_mtls(
                        certfile, keyfile, cafile, NULL /* capath - not using now */);
    
    if (context == NULL) {
        printf("Failed to create SSL context\n");
        exit(EXIT_FAILURE);
    }

    while (1) {
        int clientsock;
        // sockaddr_in -> socket address in internet style
        struct sockaddr_in clientaddr;
        unsigned int clientaddrlen = sizeof(clientaddr);

        /*
         * accept fills the size of the client address in the clientaddrlen argument
         */
        clientsock = accept(serversock, (struct sockaddr *)&clientaddr, &clientaddrlen);
        if (clientsock < 0) {
            perror("Unable to accept client connection.");
            continue;
        }

        // Create an SSL object and associate it with the context
        SSL *ssl = SSL_new(context);
        SSL_set_fd(ssl, clientsock);

        /*
         * SSL handshake done inside SSL_accept processing
         */
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
            // Once accepted, further operations are done over the SSL object.
            clientId++;
            printf("Connected to client#%d\n", clientId);

            show_cert_info(ssl);

            char message[1024] = "\0";
            recv_from_client(ssl, clientsock, &message[0], sizeof(message));

            sprintf(message, "You are client#%d", clientId);
            send_to_client(ssl, clientsock, &message[0], strlen(message));
        }
        SSL_free(ssl);
        close(clientsock);
    }

    close(serversock);
    return 0;
}