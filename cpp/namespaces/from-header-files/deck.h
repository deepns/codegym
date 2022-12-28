#include <iostream>

// TODO
// Add comments
namespace Deck
{
    typedef enum {
        CLUBS,
        SPADES,
        HEARTS,
        DIAMONDS
    } Suit;

    class Card {
    public:
        Card(int num, Suit suit) :
            _num(num),
            _suit(suit) {};

        int Num() {
            return _num;
        }

        int Suit() {
            return _suit;
        }

        void Describe() {
            std::cout << _suit << ":" << _num << std::endl;
        }

    private:
        int _num;
        int _suit;
    };
}