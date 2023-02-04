#include <vector>
#include <iostream>
#include <typeinfo>

#include "command.h"

using std::vector;

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

    auto ListNums = [](const vector<int> vec, std::string name) {
        std::cout << "Listing " << name << ", size=" << vec.size() << ", cap=" << vec.capacity() << std::endl;
        for (const auto& n : vec) {
            std::cout << n << ",";
        }
        std::cout << std::endl;
    };

    // different types of constructing vectors
    vector<int> nums{1, 2, 3, 4, 5};
    ListNums(nums, "nums");

    vector<int> nums_after_3{std::find_if(nums.begin(), nums.end(), [&](int x) { 
        return x > 3;
    }), nums.end()};
    ListNums(nums_after_3, "nums_after_3");

    // invokes the copy constructor
    vector<int> nums2 = nums;
    nums2.push_back(6);
    ListNums(nums2, "nums2");

    // invokes the move constructor
    vector<int> nums3 = std::move(nums2);
    ListNums(nums2, "nums2");
    ListNums(nums3, "nums3");

    Command c1("foo");
    std::vector<std::string> args = {"bar", "baz"};
    Command c2(args);

    args.push_back("hello");

    c1.ListArgs();
    c2.ListArgs();
        
    return 0;
};