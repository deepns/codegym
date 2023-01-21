#include <iostream>
#include <map>

using std::map;

int main()
{
    // map -sorted associative container
    // keys sorted by the provided compare function. defaults to std::less
    // backed by red-black-tree implementation underneath

    // <Key, Value, Compare, Allocator>

    // empty map
    map<std::string, int> scores;

    // iterating map
    auto list_scores = [&](std::string desc = "")
    {
        std::cout << "Listing scores: " << desc << std::endl;
        // capacity methods
        std::cout << "size=" << scores.size()
                  << " max_size=" << scores.max_size()
                  << " empty?=" << scores.empty()
                  << std::endl;

        for (const auto &[key, value] : scores)
        {
            std::cout << key << " -> " << value << std::endl;
        }
    };

    list_scores();

    // adding elements into map
    scores["adam"] = 50;

    // insert returns an iterator to the inserted element and whether
    // insertion succeeded or not
    {
        const auto [it, inserted] = scores.insert({"eve", 40});
        std::cout << "it=" << it->first
                  << " inserted=" << inserted
                  << std::endl;
    }

    auto [it, inserted] = scores.insert({"eve", 99});
    std::cout << "it=" << it->first
              << " inserted=" << inserted
              << std::endl;

    // this will overwrite though.
    scores.insert_or_assign("eve", 99);

    list_scores();

    // iterating way - 2
    std::cout << "Iterating using iterator" << std::endl;
    for (auto it = scores.begin(); it != scores.end(); ++it)
    {
        std::cout << it->first << " -> " << it->second << std::endl;
    }

    // map also supports const, reverse and const_reverse iterators

    // TODO
    // lookup methods
    //  count
    //  find
    //  contains
    //  equal range
    //  lower_bound
    //  upper_bound

    return 0;
}