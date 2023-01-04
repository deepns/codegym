#include <iostream>

class Room {
public:
    // Default constructor
    Room() {
        std::cout << "Constructing room():"
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;                  
    }

    // using member initializer
    // members are initialized before the constructor body runs
    Room(int dimension) :
        m_height(dimension), m_width(dimension), m_length(dimension) {
        std::cout << "Constructing room(dimension): "
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;
    }

    // overloaded constructors
    Room(int width, int length, int height) : 
        m_height(height), m_width(width), m_length(length) {
        std::cout << "Constructing room(height, width, length): "
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;
    }

    int area() {
        return m_width * m_height * m_length;
    }

private:
    int m_width {0};
    int m_height {0};
    int m_length {0};

};

int main()
{
    Room porch(10);
    Room dining(11, 12, 14);
    Room nospace;

    std::cout << "porch.area() = " << porch.area() << std::endl;
    std::cout << "dining.area() = " << dining.area() << std::endl;
    std::cout << "nospace.area() = " << nospace.area() << std::endl;

    return 0;
}