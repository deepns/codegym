#include <iostream>
#include <functional>
#include <list>

int main() {
    // Exploring std::list
    //  - supports constant time insertion
    //  - implemented as double linked list
    //   - fast random access is not supported

    // simple list of ints
    std::list<int> nums = {1, 2, 3, 4};

    // list of strings
    std::list<std::string> names = { "foo", "bar", "baz" };

    // capacity checks
    std::cout << "nums.empty()=" << nums.empty() << std::endl;

    std::cout << "nums.size()=" << nums.size() << " nums.max_size=" << nums.max_size() << std::endl;

    // access elements
    // access first and last elemeent
    std::cout << "names[0]=" << names.front() << std::endl;
    std::cout << "names[" << names.size() << "]=" << names.back() << std::endl;

    // iterators
    
    // iterator style - 1
    std::cout << "Iterating nums:" << std::endl;
    for (auto n : nums) {
        std::cout << n << " ";
    }
    std::cout << std::endl;

    // iterator style - 2
    std::cout << "Iterating nums:" << std::endl;
    for (const auto n : nums) {
        std::cout << n << " ";
    }
    std::cout << std::endl;

    // iterator style - 3
    std::cout << "Iterating nums:" << std::endl;
    for (auto& n : nums) {
        std::cout << n << " ";
    }
    std::cout << std::endl;

    // iterator style - 4
    std::cout << "Iterating nums:" << std::endl;
    for (auto it = nums.begin(); it != nums.end(); ++it) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    // iterator style - 5
    std::cout << "Reverse Iterating nums:" << std::endl;
    for (auto it = nums.rbegin(); it != nums.rend(); ++it) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    auto listNums = [&] () {
        std::cout << "Iterating nums (from lambda):" << std::endl;
        for (auto& n : nums) {
            std::cout << n << " ";
        }
        std::cout << std::endl;
    };

    auto listNames = [&] () {
        std::cout << "Listing names: " << std::endl;
        for (auto& name: names) {
            std::cout << name << " ";
        }
        std::cout << std::endl;
    };

    // modifiers
    // insert elements
    nums.push_back(5);
    nums.push_front(0);

    listNums();

    // clear list
    nums.clear();
    listNums();

    // inplace construction..makes difference when adding objects
    nums.emplace_front(1);
    nums.emplace_back(2);
    listNums();

    
    names.emplace_back("alice");
    listNames();

    // remove from list
    names.remove("bar");
    listNames();

    // difference between remove and erase?
    names.erase(std::find(names.begin(), names.end(),  "baz"));
    listNames();

    return 0;
}