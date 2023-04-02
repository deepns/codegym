import pytest
from happy_number import Solution

@pytest.fixture()
def happy_number_solution():
    return Solution()

@pytest.mark.parametrize("n, expected", [
    (1, True),
    (7, True),
    (10, True),
    (13, True),
    (19, True),
    (23, True),
    (2, False),
    (4, False),
    (16, False),
    (20, False),
    (89, False),
    (100, True),
    (9999999, False),
    (99887766, False),
    (123456789, False),
    (987654321, False),
    (987654329, True),
])

def test_is_happy(happy_number_solution, n, expected):
    assert happy_number_solution.isHappy(n) == expected
