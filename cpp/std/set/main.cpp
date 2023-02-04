#include <iostream>
#include <set>

// Exploring set
// data stored in an order (default comparator std::less)

int main()
{
    std::set<int> nums {1, 3, 42, 344, 254};

    // to insert
    nums.insert(5);

    // iterating a set
    // values are stored in sorted order
    std::cout << "Iterating set\n";
    for (auto& n : nums) {
        std::cout << n << std::endl;
    }

    // searching for an item
    // find() returns an iterator to the item if found
    auto it = nums.find(4234);
    if (it == nums.end()) {
        std::cerr << "number " << 4234 << " not found\n";
    }

    it = nums.find(344);
    std::cout << *it << std::endl;

    std::cout << "nums.size()=" << nums.size()
              << " nums.max_size()=" << nums.max_size()
              << std::endl;
    
    // erase all elements
    nums.erase(nums.begin(), nums.end());
    
    // construct in place
    nums.emplace(10);

    nums.insert(5);
    
    // inserting the same element again makes no difference.
    nums.insert(5);
    std::cout << "Iterating set\n";
    for (auto& n : nums) {
        std::cout << n << std::endl;
    }
    return 0;
}