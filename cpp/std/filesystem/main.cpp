#include <filesystem>
#include <iostream>

// exploring filesystem functions
//
int main()
{
    // Get current path
    // std::filesystem::current_path returns the path of the program
    auto cur_path = std::filesystem::current_path();
    std::cout << "cur_path=" << cur_path << std::endl;

    // the same can be used to change the path
    std::filesystem::current_path(cur_path.parent_path());
    std::cout << "cur_path=" << std::filesystem::current_path() << std::endl;

    // Iterating a directory
    for (auto const &dir_entry : std::filesystem::directory_iterator{cur_path})
    {
        std::cout << dir_entry.path() << '\n';
    }

    std::cout << "cur_path.filename()=" << cur_path.filename()
              << " cur_path.relative_path()=" << cur_path.relative_path()
              << " cur_path.stem()=" << cur_path.stem()
              << std::endl;

    // also supports .. based navigation
    // can use just like path on the shell
    cur_path = cur_path / "..";
    std::cout << "cur_path=" << std::filesystem::current_path() << std::endl;

    // iterating a directory recursively
    for (auto const &dir_entry :
         std::filesystem::recursive_directory_iterator{cur_path})
    {
        // fsize doesn't work on directories
        if (dir_entry.is_regular_file())
        {
            std::cout << "fname=" << dir_entry.path()
                      << " fsize=" << dir_entry.file_size()
                      << std::endl;
        }
    }

    return 0;
}