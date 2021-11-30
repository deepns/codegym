def right_rotate(lst, k):
    # Implement a function right_rotate(lst, k) which will rotate the given list by k.
    # This means that the right-most elements will appear at the left-most position in
    # the list and so on. You only have to rotate the list by one element at a time.

    # sample input:
    # lst = [10,20,30,40,50]
    # k = 3

    # output = [30, 40, 50, 10, 20]

    # notes:
    # k - constraints? less than length of the lst or can be any arbitrary value?
    #   - negative?
    # rotate one element at a time swap(lst[0], lst[:-1]) - no.
    #   copy lst[-1] to temp
    #   move each element to its next index
    #   for i in range(len(lst)-1, -1, 0):
    #       lst[i] = lst[i-1]
    #   lst[0] = temp

    # this has to be done k number of times.
    # thats O(nk)
    # if k == n, it will be O(n^2)

    # for k >= n -> number of rotations would be n mod k

    # lets say k > 0.

    # corner cases
    if not lst or k <= 0:
        return lst
    
    k = k % len(lst)

    # for in place rotation, manually swap the elements
    # for _ in range(k):
    #     tmp = lst[-1]
    #     for i in range(N-1, 0, -1):
    #         lst[i] = lst[i-1]
    #     lst[0] = tmp
    
    # return lst

    # to rotate right, bring k elements from the right to front.
    # move n-k elements from the left to the right.
    # but this would be moving more than one item at a time though.
    # more pythonic way to do than manual rotation
    lst = lst[-k:] + lst[:-k]
    return lst

def test_right_rotate():
    testcases = [
        [[[1, 2, 3, 4, 5], 2], [4, 5, 1, 2, 3]],
        [[[], 1], []],
        [[[300, -1, 3, 0], 3], [-1, 3, 0, 300]],
        [[['right', 'rotate', 'python'], 4], ['python', 'right', 'rotate']],
        [[[], 10], []]
    ]

    for testcase in testcases:
        inp, expected_output = testcase
        lst, k = inp[0], inp[1]
        actual_output = right_rotate(lst, k)
        assert expected_output == actual_output

test_right_rotate()