#include <iostream>
#include <typeinfo>

int main()
{
    std::cout << "=====> Learning typeid" << std::endl;
    auto name = "foo";
    auto age = 25;
    auto ptr_age = &age;
    auto score = 98.2;
    std::string fullname("Foo, Bar");

    // typeid operator - to find the runtime type of an object
    // typeid(expression) returns std::type_info object
    // type_info::name() returns a compiler dependent byte string to identify a
    // type PKc for string literals
    std::cout << "typeid(name)=" << typeid(name).name() << std::endl;
    // i for integer, Pi for pointer to integer
    std::cout << "typeid(age)=" << typeid(age).name() << std::endl;
    std::cout << "typeid(ptr_age)=" << typeid(ptr_age).name() << std::endl;

    // d for double
    std::cout << "typeid(score)=" << typeid(score).name() << std::endl;

    // something like this for std::string
    // =NSt7__cxx1112basic_stringIcSt11char_traitsIcESaIcEEE
    std::cout << "typeid(fullname)=" << typeid(fullname).name() << std::endl;

    return 0;
}