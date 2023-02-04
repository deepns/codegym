#include <iostream>
#include <iterator>
#include <list>

int main() {
    // Exploring std::list
    //  - supports constant time insertion
    //  - implemented as double linked list
    //   - fast random access is not supported

    // simple list of ints
    std::list<int> nums = {1, 2, 3, 4};

    // list of strings
    std::list<std::string> names = { "foo", "bar", "baz", "zebra", "cat" };

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

    auto listNames = [&] (std::string desc="") {
        std::cout << "Listing names: " << desc << std::endl;
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
    // remove() returns the number of elements removed (only from c++20)
    auto numNamesRemoved = names.remove("bar");
    std::cout << "numNamesRemoved=" << numNamesRemoved << std::endl;
    listNames();

    // no error or exception when removing an non-existing element
    numNamesRemoved = names.remove("bar");
    std::cout << "numNamesRemoved=" << numNamesRemoved << std::endl;
    listNames("after remove bar again");

    // remove_if
    names.remove_if([](const std::string& name) {
        return name.starts_with("c");
    });
    listNames("after remove_if");

    // erase from the list
    names.erase(std::find(names.begin(), names.end(),  "baz"));
    listNames("after erase:");
    
    // difference between remove and erase?
    // erase takes the iterator pos as input, and can also take 
    // range of elements
    nums = {1, 10, 23, 24, 12, 84, 24};
    auto start = nums.begin();
    auto end = nums.end();
    std::advance(start, 2);
    std::advance(end, -2);

    listNums();

    std::cout << "Erasing these elements..." << std::endl;
    for (auto it=start; it != end; ++it) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    nums.erase(start, end);
    listNums();
    
    return 0;
}