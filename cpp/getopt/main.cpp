#include <iostream>
#include <getopt.h>
#include <unistd.h>

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
        // TODO update the help description
        std::cout << "./main -r/--required ";
    }

    void ParseArgsWithGetOpt(int argc, char **argv)
    {
        // getopt can parse three types of options
        //  - an option that require argument   // specified by option letter followed by :
        //  - an option that takes an optional argument // specified by option letter followed by ::
        //  - an option that doesn't need any argument // specified by just the option letter
        const char *short_options = "r:o::h";

        // argument for an option may or may not be separated by spaces.
        // both are valid here.
        // -o val
        // -oval

        // getopt returns -1 when it has no more arguments to process
        int ch = -1;
        while ((ch = getopt(argc, argv, short_options)) != -1)
        {
            switch (ch)
            {
            case 'r':
                std::cout << "r=" << optarg << std::endl;
                break;
            case 'o':
                // There is a gotcha here. If the optional argument and its value are
                // separated by space, then getopt cannot differentiate whether the
                // next token in the command option is another argument or value to
                // the preceding argument. In this example, with "-o someval", optarg
                // will be NULL when parsing the option 'o'.
                //
                // man page for getopt says this
                // On return from getopt(), optarg points to an option argument, if it
                // is anticipated, and the variable optind contains the index to the next
                // argv argument for a subsequent call to getopt().  The variable optopt
                // saves the last known option character returned by getopt().

                // for optional arguments where argument and value are separated by space,
                // we can look ahead into argv using optind index since optind is set to
                // next argv argument for the subsequent getopt().

                std::cout << "o="
                          << (optarg ? optarg : "NULL")
                          << std::endl;
                if (OPTIONAL_ARGUMENT_IS_PRESENT)
                {
                    std::cout << "o=" << optarg << std::endl;
                }
                break;
            case '?':
                // returned when an option not specified in the short_options is encountered
                // an error message is written to stderr as well (can be disabled by setting
                // opterr to 0)
            case 'h':
                ShowHelp();
            default:
                break;
            }
        }
    }

    void ParseArgsWithGetOptLong(int argc, char **argv)
    {
        // TODO
        // Clean up the comments here.
        
        option long_options[] = {
            {"required", required_argument, NULL, 'r'},
            {"optional", optional_argument, NULL, 'o'},
            {"help", no_argument, NULL, 'h'},
            {nullptr, 0, nullptr, 0}};

        // getopt can parse three types of options
        //  - an option that require argument   // specified by option letter followed by :
        //  - an option that takes an optional argument // specified by option letter followed by ::
        //  - an option that doesn't need any argument // specified by just the option letter
        const char *short_options = "r:o::h";

        // with short options,
        // -o val
        // -oval

        // With the long options, it supports two more ways of specifying the options.
        //
        // --option=val
        // --option val

        // getopt_long returns -1 when it has no more arguments to process
        int ch = -1;
        while ((ch = getopt_long(argc, argv, short_options, long_options, nullptr)) != -1)
        {
            switch (ch)
            {
            case 'r':
                std::cout << "r=" << optarg << std::endl;
                break;
            case 'o':
                // There is a gotcha here. If the optional argument and its value are
                // separated by space, then getopt cannot differentiate whether the
                // next token in the command option is another argument or value to
                // the preceding argument. In this example, with "-o someval", optarg
                // will be NULL when parsing the option 'o'.
                //
                // man page for getopt says this
                // On return from getopt(), optarg points to an option argument, if it
                // is anticipated, and the variable optind contains the index to the next
                // argv argument for a subsequent call to getopt().  The variable optopt
                // saves the last known option character returned by getopt().

                // for optional arguments where argument and value are separated by space,
                // we can look ahead into argv using optind index since optind is set to
                // next argv argument for the subsequent getopt().

                std::cout << "o="
                          << (optarg ? optarg : "NULL")
                          << std::endl;
                if (OPTIONAL_ARGUMENT_IS_PRESENT)
                {
                    std::cout << "o=" << optarg << std::endl;
                }
                break;
            case '?':
                // returned when an option not specified in the short_options is encountered
                // an error message is written to stderr as well (can be disabled by setting
                // opterr to 0)
            case 'h':
                ShowHelp();
            default:
                break;
            }
        }
    }

} // namespace

int main(int argc, char **argv)
{
    ParseArgsWithGetOptLong(argc, argv);
}