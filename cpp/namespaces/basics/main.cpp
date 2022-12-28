/*
 * Getting my feet wet with namespaces
 */

#include <iostream>

// simple namespace with a function and datatype inside
namespace Foo {
    class Bar {
    public:
        Bar(int id) : _id(id) {};
        int Id() {
            return _id;
        }

    private:
        int _id;
    };

    void hello(std::string name) {
        std::cout << "Hello, " << name << std::endl;
    }

    void echo(int num) {
        std::cout << "Your number is, " << num << std::endl;
    }
}

// namespace can be declared in multiple places within a single file or multiple files
namespace Foo {
    void echo(std::string name, int count) {
        for (int i=0; i < count; i++) {
            std::cout << name << "..." << std::endl;
        }
    }
}

int main() {
    // Access namespace members
    Foo::hello("Bar");
    Foo::echo(Foo::Bar(101).Id());
    Foo::echo("bark", 3);

    return 0;
}