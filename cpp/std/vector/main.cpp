#include <vector>
#include <iostream>
#include <typeinfo>

int main() {
    // Exploring std::vector

    // declaration
    std::vector<std::string> names = {
        "foo",
        "bar",
        "baz"
    };

    auto ListNames = [&] () {
        std::cout << "Iterating names..." << std::endl;
        for (auto& name : names) {
            std::cout << name << std::endl;
        }
    };

    ListNames();

    // capacity checks
    std::cout << "names.capacity() = " << names.capacity() << std::endl;
    std::cout << "names.size() =" << names.size() << std::endl;
    std::cout << "names.max_size() = " << names.max_size() << std::endl;
    
    // access elements
    
    // direct access
    std::cout << "names.at(1) = " << names.at(1) << std::endl;
    std::cout << "names.front() = " << names.front() << std::endl;
    std::cout << "names.back() = " << names.back() << std::endl;
    std::cout << "using [] operator: names[0] = " << names[0] << std::endl;
 
    // iterators

    // using range
    for (const auto& n : names) {
        std::cout << n << ", typeid = " << typeid(n).name() << std::endl;
    }

    // Does this make a new copy of the vector member?
    for (auto n = names.begin(); n != names.end(); ++n) {
        std::cout << "n = " << *n << std::endl;
    }

    for (auto n = names.cbegin(); n != names.cend(); ++n) {
        std::cout << "n = " << *n << std::endl;
    }

    // c-style iteration
    for (auto i = 0; i < names.size(); i++) {
        std::cout << "names[" << i << "] = " << names[i] << std::endl;
    }

    // modifiers

    // clear vector
    names.clear();

    // push new element..can only push to the back
    names.push_back("banana");

    // insert based on the iterator position
    names.insert(names.begin(), "apple");

    std::cout << "popping from names..." << std::endl;
 
    ListNames();

    return 0;
};