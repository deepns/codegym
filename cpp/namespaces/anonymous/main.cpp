#include <iostream>
#include <unistd.h>

// Anonymous namespaces
// - To ensure names have internal linkage
// - restricts the scope of the symbols to the file
// - Gets the same behavior as definining global variables with static scope
//     - So names under an anonymous namespace accessible only within the file
// - Defining an anonymous namespaces in a header file is not of use, as each file
//   including the header will get its own copy of the anonymous namespace
namespace {
    int echo_count = 1;
    void echo(std::string text) {
        std::cout << "Echoing..." << echo_count << " time(s)" << std::endl;
        for (int i = 0; i < echo_count; i++) {
            std::cout << text << std::endl;
            sleep(1);
        }
    }
}

int main() {
    // names in anonymous namespace are accessed just like global variables
    echo("Foo");

    // set echo_count to a random variable
    srand(time(NULL));
    echo_count = rand() % 10;

    echo("Bar");
    return 0;
}
