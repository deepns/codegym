// A sample program to learn getopt_long

#include <iostream>
#include <getopt.h>
#include <stdlib.h>

// Thanks to https://cfengine.com/blog/2021/optional-arguments-with-getopt-long/
// for this macro
#define OPTIONAL_ARGUMENT_IS_PRESENT                             \
    ((optarg == NULL && optind < argc && argv[optind][0] != '-') \
         ? (bool)(optarg = argv[optind++])                       \
         : (optarg != NULL))

namespace
{
    void ShowHelp()
    {
        std::cout << "./si -p/--principal <value> -n/--num-of-years <value> [-r/--rate <value>] [-h/--help]" << std::endl;
    }

    void ParseOptions(int argc, char *argv[], double &p, double &n, double &r)
    {
        option long_options[] = {
            {"principal", required_argument, nullptr, 'p'},
            {"num-of-years", required_argument, nullptr, 'n'},
            {"rate", optional_argument, nullptr, 'r'},
            {"help", no_argument, nullptr, 'h'},
            {nullptr, 0, nullptr, 0}};

        const char *short_options = "p:n:r::h";

        int ch = -1;
        while ((ch = getopt_long(argc, argv, short_options, long_options, nullptr)) != -1)
        {
            switch (ch)
            {
            case 'p':
                p = std::stod(optarg);
                break;
            case 'n':
                n = std::stod(optarg);
                break;
            case 'r':
                if (OPTIONAL_ARGUMENT_IS_PRESENT)
                {
                    r = std::stod(optarg) / 100;
                    if (r < 0 or r > 100)
                    {
                        std::cerr << "You kidding??..Rate must be between 0..100" << std::endl;
                        exit(1);
                    }
                }
                break;
            case '?':
            case 'h':
            default:
                ShowHelp();
                exit(1);
                break;
            }
        }
    }
} // namespace

int main(int argc, char *argv[])
{
    double principal = 0;
    double num_of_years = 1;
    double rate = .05;

    ParseOptions(argc, argv, principal, num_of_years, rate);

    double si = principal * rate * num_of_years;
    std::cout << "Simple interest = " << si << std::endl;
    std::cout << "Principal + interest = " << principal + si << std::endl;
    return 0;
}