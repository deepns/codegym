#include <iostream>

int main()
{
    // bare simple lambda function
    auto max = [](int a, int b) {
        if (a > b) {
            return a;
        } else {
            return b;
        }
    };

    auto pi = 3.14;

    // another lambda that uses the variables from the outer scope
    auto area = [&](int r) {
        return pi * r * r;
    };

    std::cout << "Calling max lambda: max(10, 20): " << max(10, 20) << std::endl;
    std::cout << "Calling area lambda: area(5): " << area(5) << std::endl;
}