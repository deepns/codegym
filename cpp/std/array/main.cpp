#include <iostream>
#include <array>

int main()
{
    // Fixed size array of a certain type
    // What is the advantage of this over old style?
    // say std::string choices[2] in this case?
    //
    // when an array would be used?
    std::array<std::string, 4> choices = {"Yes", "No", "May be"};

    for (auto &choice : choices)
    {
        std::cout << choice << std::endl;
    }

    // I thought this would error out.. but it didn't
    // no compile time or runtime fatal error.
    // caused memory corruption? need to check
    choices[4] = "No comments";

    // exploring the member functions
    std::cout << "choices.at(0)=" << choices.at(0) << std::endl
              << "choices.front()=" << choices.front() << std::endl
              << "choices.back()=" << choices.back() << std::endl
              << "choices.empty()=" << (choices.empty() ? "Empty" : "Not Empty") << std::endl
              << "choices.max_size()=" << choices.max_size() << std::endl
              << "choices.size()=" << choices.size() << std::endl // size and max_size() are same
              << "choices[2]=" << choices[2] << std::endl; // index based lookup
}