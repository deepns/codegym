#include <iostream>

class Room
{
public:
    // Default constructor
    Room()
    {
        m_width = m_length = m_height = 0;
        std::cout << "Constructing room():"
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;
    }

    // using member initializer
    // members are initialized before the constructor body runs
    // can members be initialized in any order?
    Room(int dimension) : m_height(dimension), m_width(dimension), m_length(dimension)
    {
        std::cout << "Constructing room(dimension): "
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;
    }

    // overloaded constructors
    Room(int width, int length, int height) : m_height(height), m_length(length), m_width(width)
    {
        std::cout << "Constructing room(height, width, length): "
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;
    }

    int area()
    {
        return m_width * m_height * m_length;
    }

private:
    int m_width;
    int m_height;
    int m_length;
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