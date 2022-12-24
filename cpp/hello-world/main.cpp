#include <iostream>

int main (int argc, char *argv[]) {
    auto name = "FooBar";

    if (argc > 1) {
        name = argv[1];
    }

    std::cout << "Hello, " << name << std::endl;
    return 0;
}