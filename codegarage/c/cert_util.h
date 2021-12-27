#include <openssl/x509.h>

void show_cert_info(X509 *cert);
void show_cert_info_from_file(const char *certfilepath);