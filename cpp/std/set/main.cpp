#include <iostream>
#include <random>
#include <set>

// Exploring set
// data stored in an order (default comparator std::less)

std::set<int> generateRandomIntegers(int from = 1, int to = 100, int count = 10)
{
    std::set<int> result;
    std::random_device rd;

    // std::mt19937 implements a random number engine based on the
    // Mersenne Twister algorithm. It generates a large number of random
    // numbers with a long period, making it suitable for applications that
    // require a large amount of random numbers. The algorithm is also
    // deterministic, meaning that it generates the same sequence of
    // random numbers every time it is seeded with the same value
    std::mt19937 gen(rd());

    // can generate random numbers using the operator() function of mt19937.
    // std::cout << gen() << std::endl;
    // or tie it with a specific distribution

    // std::uniform_int_distribution defines a uniform distribution over a
    // range of integers, meaning that each integer in the range has an
    // equal probability of being generated. 
    // <> let the compiler deduce the template type
    std::uniform_int_distribution<> dis(from, to);

    for (int i = 0; i < count; ++i)
    {
        result.insert(dis(gen));
    }

    return result;
}

int main()
{
    std::set<int> nums {1, 3, 42, 344, 8, 10, 11, 19, 34, 254};

    // to insert
    nums.insert(5);

    // iterating a set
    // values are stored in sorted order
    auto listNums = [&nums](const std::string& desc="") {
        std::cout << desc << std::endl;
        for (auto& n : nums) {
            std::cout << "[ " << n << " ] " << std::endl;
        }
    };

    listNums("Iterating set");
    
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
    
    // erase first 3 elements
    nums.erase(nums.begin(), std::next(nums.begin(), 3));
    listNums("After erasing first 3 elements");

    // erase last two elements
    nums.erase(std::prev(nums.end(), 2), nums.end());
    listNums("After erasing last 2 elements");

    // erase all elements
    nums.erase(nums.begin(), nums.end());
    listNums("After erasing all elments");
    
    // construct in place
    nums.emplace(10);

    nums.insert(5);
    
    // inserting the same element again makes no difference.
    nums.insert(5);
    listNums("After inserting same element twice");

    nums = generateRandomIntegers();
    listNums("After generating some random numbers");

    it = nums.lower_bound(40);
    while (it != nums.end()) {
        std::cout << *it << std::endl;
        ++it;
    }

    return 0;
}