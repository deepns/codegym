#include <iostream>
#include <unistd.h>

void show_help() {
    std::cout << "./a.out -s <server> -p [-h]";
}

int main(int argc, char **argv) {
    // : - require argument
    // :: - optional argument?
    const char *opts = "s:ph";
    int ch;
    while ((ch = getopt(argc, argv, opts)) != -1) {
        switch (ch)
        {
        case 's':
            std::cout << "ch=" << optarg << std::endl;
            break;
        case 'p':
            std::cout << "ch=" << (optarg ? optarg : "NULL") << std::endl;
            break;
        case '?': // for unrecognized options
        case 'h': // for help
        default:
            show_help();
            break;
        }
    }
    return 0;
}