#include <iostream>

class Noconstruct {

public:
    // Prevent compiler from generating default constructor for this class
    // this marks the constructor function as deleted
    // what is the use of such non-constructible class?
    Noconstruct() = delete;
};

class Room {
public:
    // Default constructor
    Room() {
        std::cout << "Constructing room():"
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;                  
    }

private:
    int m_width {0};
    int m_height {0};
    int m_length {0};

};

int main() {
    // compilation error: default constructor of class can't be referenced
    // Noconstruct n1;
    Room trial;

    return 0;
}