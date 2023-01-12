#include <iostream>
#include <typeinfo>
#include <array>

// Learning about lambdas
//
// Syntax:
//
//  [ capture clause ] ( parameters ) -> return_type {
//      // function body
// }

int main()
{
    // bare simple lambda function
    auto max = [](int a, int b)
    {
        if (a > b)
        {
            return a;
        }
        else
        {
            return b;
        }
    };

    auto pi = 3.14;

    // another lambda that uses the variables from the outer scope
    auto area = [&](int r)
    {
        return pi * r * r;
    };

    std::cout << "Calling max lambda: max(10, 20): " << max(10, 20) << std::endl;
    std::cout << "Calling area lambda: area(5): " << area(5) << std::endl;

    // type(num) in this case is initializer list
    auto nums = {1, 3, 4, 9, 12, 14, 15};
    std::cout << "typeid(nums) = " << typeid(nums).name() << "\n";

    // same as above.
    auto colors = {"orange", "black", "green"};
    std::cout << "typeid(colors) = " << typeid(colors).name() << "\n";

    auto has_green = std::find(colors.begin(), colors.end(), "green");
    if (has_green == colors.end()) {
        std::cout << "Can't find green\n";
    } else {
        std::cout << "Found green!\n";
    }

        // find_if returs an iterator
    // pass a lambda function to the find_if predicate
    auto has_black = std::find_if(colors.begin(), colors.end(), [](std::string color)
                                  { return color == "black"; });
    std::cout << "has_black=" << has_black << " typeid(has_black)=" << typeid(has_black).name() << "\n";

    if (has_black == colors.end())
    {
        std::cout << "Can't find black\n";
    }
    else
    {
        std::cout << "Found black!\n";
    }

    std::array<int, 10> ages{3, 4, 5, 2, 3, 4, 5, 2, 7, 8};
    auto ages_below_5 = std::find_if(ages.begin(), ages.end(), [](int age)
                                     { return age < 5; });

    if (ages_below_5 != ages.end()) {
        std::cout << "Has ages below 5" << std::endl;
    }
    return 0;
}