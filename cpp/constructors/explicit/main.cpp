#include <iostream>
#include <cstdio>

// cstdio imports the stdio symbols into the standard namespace
// including stdio.h directly can cause collision with the symbols
// in the global namespace. not a c++ way of doing things.

class Room
{
public:
    // If a constructor has a single parameter, it allows implicit conversion of
    // the parameter type to the class type
    // Room r = 10;
    //
    explicit Room(int dimension) : m_height(dimension), m_width(dimension), m_length(dimension)
    {
        std::cout << "Constructing room(dimension): "
                  << "width=" << m_width << " "
                  << "height=" << m_height << " "
                  << "length=" << m_length << std::endl;
    }

    void show_dimensions()
    {
        printf("Room: width=%d, height=%d, length=%d\n", m_width, m_height, m_length);
    }

    int area()
    {
        return m_width * m_height * m_length;
    }

    void set_height(int height)
    {
        m_height = height;
    }

    void set_width(int width)
    {
        m_width = width;
    }

    void set_length(int length)
    {
        m_length = length;
    }

    int width() { return m_width; }
    int length() { return m_length; }
    int height() { return m_height; }

private:
    int m_width;
    int m_height;
    int m_length;
};

void expand(Room room, int feet)
{
    if (feet > 0)
    {
        room.set_width(room.width() + feet);
        room.set_length(room.length() + feet);
        room.set_height(room.height() + feet);
    }
    std::cout << "New dimensions of room:" << std::endl;
    room.show_dimensions();
}

int main()
{
    // This type of initialization is made possible by the constructors that take single parameter
    // With explicit keyword in place, this is not allowed.
    // shows the erorr:
    // no suitable constructor exists to convert from "int" to "Room"

    // Room one = 10; // 10x10x10 by room
    // This is allowed though
    Room one(10);

    // implicit type conversion happen at the parameter level too.
    // here passing "20" in the first parameter converts 20 into the
    // actual parameter type (Room)
    // throws the same compilation error as above
    // expand(20, 2);
    expand(one, 12);

    // Such usage can often lead to confusion and errors in the code
    // Making the constructor explicit prevents such implicit conversion

    return 0;
}
