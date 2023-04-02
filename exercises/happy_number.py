# https://leetcode.com/problems/happy-number/description/
#
# Write an algorithm to determine if a number n is happy.

# A happy number is a number defined by the following process:

# Starting with any positive integer, replace the number by the sum of the squares of its digits.
# Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
# Those numbers for which this process ends in 1 are happy.
# Return true if n is a happy number, and false if not.

def sum_of_squares_of_digits(n):
    def square_of_digits(n):
        while n > 0:
            yield (n % 10)**2
            # n / 10 will result in a float
            # n // 10 will result in an int
            n = n // 10
    return sum(square_of_digits(n))            

class Solution:
    def isHappy(self, n: int) -> bool:
        # what is a happy number?
        # n > 0
        # n = sum(squares_of_digits_of_n) 
        # if n reaches 1, n is happy
        # if not (i.e. stuck in endless cycle), then n is not happy

        # find the sum..capture in a set
        # if the new sum == 1, n is happy
        # if new sum is in the set, we have a cycle
        # but that requires additional space
        # how to find the cycle?

        # progressing linearly here
        # next step determined by sum_of_digits
        # use slow..fast method to detect cycle

        slow = n
        fast = sum_of_squares_of_digits(n)
        while slow != fast and fast != 1:
            slow = sum_of_squares_of_digits(slow)
            fast = sum_of_squares_of_digits(sum_of_squares_of_digits(fast))
        
        return True if fast == 1 else False

