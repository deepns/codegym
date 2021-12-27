#include <stdio.h>

#include "cert_util.h"

int main(int argc, char **argv)
{
    if (argc < 2) {
        fprintf(stderr, "Error: Cert file is missing\n");
        exit(EXIT_FAILURE);
    }

    show_cert_info_from_file(argv[1]);
    return 0;
}