#include <iostream>
#include <random>
#include <set>

// Thank you ChatGPT

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
    std::set random_ints = generateRandomIntegers();

    for (auto &n : random_ints)
    {
        std::cout << n << std::endl;
    }
}