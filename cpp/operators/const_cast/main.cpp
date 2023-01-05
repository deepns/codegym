#include <iostream>

namespace
{
    void echo_c_string(char *str)
    {
        std::cout << str << std::endl;
    }

    void echo_string(std::string str)
    {
        std::cout << str << std::endl;
    }

} // namespace

int main()
{
    std::cout << "=====> Learning const_cast operator" << std::endl;
    const int a = 10;
    const int *ptr_a1 = &a;
    std::cout << "a=" << a << ","
              << "ptr_a1=" << *ptr_a1 << std::endl;

    // error: read-only variable is not assignable
    // *ptr_a = 11;

    // Can change the value pointed to by the const pointer
    // by casting the address to a non-const pointer
    int *ptr_a2 = const_cast<int *>(ptr_a1);

    // This doesn't change the original const variable "a"
    *ptr_a2 = 20;
    std::cout << "a=" << a << ","
              << "ptr_a1=" << *ptr_a1 << std::endl;

    // When is it useful?
    // Some C functions take non-const arguments but not necessarily make any
    // modifications to the arguments.
    // for e.g. string::c_str() returns a const pointer to the c-string
    // representation of the string. Functions taking only 'char *' cannot take
    // the values returned by c_str() directly.
    std::string name = "foo";
    echo_c_string(const_cast<char *>(name.c_str()));

    // raises the compiler warning "ISO C++11 does not allow conversion from
    // string literal to 'char *"
    // echo_c_string("bar");

    // this works fine though.
    echo_string("bar");

    return 0;
}