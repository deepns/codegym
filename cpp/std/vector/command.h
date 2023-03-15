#include <iostream>
#include <vector>

class Command {
private:
    std::vector<std::string> args_;
public:
    Command(const std::string& arg) : args_({arg}) {}
    Command(const std::vector<std::string>& args) : args_(args) {}
    void ListArgs() const {
        for (auto& arg : args_) {
            std::cout << "arg=" << arg << std::endl;
        }
    };
};