#include <filesystem>
#include <iostream>

// exploring filesystem functions
//
int main() {
    // Get current path
    // std::filesystem::current_path returns the path of the program
    auto cur_path = std::filesystem::current_path();
    std::cout << "cur_path=" << cur_path << std::endl;

    // the same can be used to change the path
    std::filesystem::current_path(cur_path.parent_path());
    std::cout << "cur_path=" <<  std::filesystem::current_path() << std::endl;

    // Iterating a directory
    for (auto const& dir_entry : std::filesystem::directory_iterator{cur_path}) {
        std::cout << dir_entry.path() << '\n';
    }
    return 0;
}