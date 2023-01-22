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

    auto list_map = [&](map<std::string, int> a_map, std::string desc = "")
    {
        std::cout << "Listing map: " << desc << std::endl;
        // capacity methods
        std::cout << "size=" << a_map.size()
                  << " max_size=" << a_map.max_size()
                  << " empty?=" << a_map.empty()
                  << std::endl;

        for (const auto &[key, value] : a_map)
        {
            std::cout << key << " -> " << value << std::endl;
        }
    };

    list_map(scores);

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

    list_map(scores, "after insertion");

    // iterating way - 2
    std::cout << "Iterating using iterator" << std::endl;
    for (auto it = scores.begin(); it != scores.end(); ++it)
    {
        std::cout << it->first << " -> " << it->second << std::endl;
    }

    // insert some more
    auto result = scores.insert({"bob", 100});
    it = result.first;
    inserted = result.second;

    // with bob inserted, will advancing the iterator take to the next element in the order?
    // Yes.. the iterator does continue in order.
    std::cout << "it=" << it->first << " ++it=" << (++it)->first << std::endl;
    // map also supports const, reverse and const_reverse iterators

    // lookup methods

    // contains is from c++20 onwards
    if (scores.contains("bob")) {
        std::cout << "scores[\"bob\"]=" << scores["bob"] << std::endl;
    }

    // what happens when looking up an non-existent key using [] operator?
    auto _x = scores["a-missing-key"];
    std::cout << "_x=" << _x << std::endl;

    // it performs an insertion with the default value (0 in this case)
    // can see that the key 'a-missing-key' is inserted into the map
    list_map(scores);

    // extract - takes the element out of the container and returns it in
    // a node_handle (https://en.cppreference.com/w/cpp/container/node_handle)
    auto nh_bob = scores.extract("bob");
    list_map(scores);

    std::cout << "nh_bob.key()=" << nh_bob.key() // key() returns a non-const reference to the key
              << " nh_bob.mapped()=" << nh_bob.mapped() // returns a const reference to the mapped value
              << std::endl;

    nh_bob.key() = "bob-galley";
    
    // scores.insert(nh_bob) didn't work. failed to compile.
    // scores.insert(std::move(nh_bob)) worked. What std::move did here?
    // this results in segfault. why?
    // can't seem to be able to insert into the same map.
    // trying another one
    // map<std::string, int> scores2;
    // const auto status = scores2.insert(std::move(nh_bob));

    // just moving didn't cause a segfault. then what else?
    // std::move(nh_bob);
    auto x = std::move(nh_bob);
    std::cout << x.key() << " -> " << x.mapped() << " " << x.empty() << std::endl;

    // this overloaded version of insert() returns an insert_return_type
    // https://en.cppreference.com/w/cpp/container/map#Member_types
    // template<class Iter, class NodeType>
    // struct /*unspecified*/
    // {
    //     Iter     position;
    //     bool     inserted;
    //     NodeType node;
    // };

    // std::cout << "status.position->first=" << status.position->first
    //           << " status.position->second=" << status.position->second
    //           << " status.node.key()=" << status.node.key()
    //           << " status.node.mapped()=" << status.node.mapped()
    //           << std::boolalpha << " status.inserted=" << status.inserted
    //           << std::noboolalpha
    //           << std::endl;

    // some more search methods
    //  equal range
    //  lower_bound
    //  upper_bound

    return 0;
}