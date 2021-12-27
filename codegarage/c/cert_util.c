#include <openssl/pem.h>
#include <openssl/x509.h>
#include <stdio.h>

#include "cert_util.h"

void show_cert_info(X509 *cert)
{
    if (!cert) {
        fprintf(stderr, "Invalid cert\n");
    }

    printf("================================================\n");

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
    
    // No _fp functions available with ASN1_TIME, so using BIO interfaces.
    // A good stackoverflow discussion on BIO interfaces
    // https://stackoverflow.com/questions/51672133/what-are-openssl-bios-how-do-they-work-how-are-bios-used-in-openssl

    BIO *bio_out;
    bio_out = BIO_new_fp(stdout, BIO_NOCLOSE);

    X509_NAME *issuer = X509_get_issuer_name(cert);
    printf("Cert issued by:\n");
    X509_NAME_print_ex_fp(stdout, issuer, 2 /*indent*/, XN_FLAG_ONELINE);
    printf("\n");

    X509_NAME *subject = X509_get_subject_name(cert);
    printf("Cert subject:\n");
    X509_NAME_print_ex_fp(stdout, subject, 2 /*indent*/, XN_FLAG_MULTILINE);
    printf("\n");

    printf("Validity:\n");
    printf("\tCert not valid before: ");
    const ASN1_TIME *not_before = X509_get0_notBefore(cert);
    ASN1_TIME_print(bio_out, not_before);
    printf("\n");

    printf("\tCert not valid after: ");
    const ASN1_TIME *not_after = X509_get0_notAfter(cert);
    ASN1_TIME_print(bio_out, not_after);
    printf("\n");

    // Get the public key and print some info about the key.

    // X509_get_pubkey bumps up the refcount on the public key
    // in the cert. caller must call EVP_PKEY_free on that.
    EVP_PKEY *pubkey = X509_get_pubkey(cert);

    // Since OpenSSL 1.1, all structures in libssl public headers
    // have been removed. Callers must use the accessor funtions to
    // to access the associated values. See https://wiki.openssl.org/index.php/OpenSSL_1.1.0_Changes
    int type = EVP_PKEY_base_id(pubkey);
    switch(type) {
        case EVP_PKEY_RSA:
        {
            RSA *rsa = EVP_PKEY_get0_RSA(pubkey);
            if (rsa) {
                RSA_print(bio_out, rsa, 0 /*offset*/);
                // other option: RSA_print_fp(stdout, rsa, 0);
                const RSA_METHOD *rsa_method = RSA_get_method(rsa);
                const char *method_name = RSA_meth_get0_name(rsa_method);
                printf("RSA method=%s\n", method_name);        
            }
        }
        break;
        case EVP_PKEY_EC: // Elliptic Curve Key
        {
            EC_KEY *eckey = EVP_PKEY_get0_EC_KEY(pubkey);
            if (eckey) {
                EC_KEY_print(bio_out, eckey, 0);
            }
        }
        break;
        default:
            printf("Unsupported key type %d\n", type);

    }
    EVP_PKEY_free(pubkey);

    BIO_free(bio_out);
    X509_free(cert);
    printf("================================================\n");
}

void show_cert_info_from_file(const char *certfilepath)
{
    FILE *certfile = fopen(certfilepath, "r");

    if (!certfile) {
        perror("failed to read cert");
    }

    X509 *cert = PEM_read_X509(certfile, NULL, NULL, NULL);

    if (!cert) {
        fprintf(stderr, "Failed to read cert from %s\n", certfilepath);
    }

    show_cert_info(cert);
    X509_free(cert);
    fclose(certfile);
}