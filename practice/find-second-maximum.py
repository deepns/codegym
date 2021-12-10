def find_second_maximum(lst):
    """
    Returns the second largest element in the list.
    """
    # sample input: [9,2,3,6], output: 6

    # notes:
    # input must have at least two elements in the list
    # can there be duplicates?
    # for e.g. [2, 3, 4, 4] -> output 3 or 4?
    # but for [4, 4]? => input error.

    # largest -> sort by acsending order
    # second largest -> return the second from the last in the sorted list
    # sorting would be nlogn. can we do better? without sorting?

    # or sort in reverse order and return the second item from the beginning

    # for i, x in enumerate(sorted(lst, reverse=True)):
    #     if i == 1:
    #         return x

    # Trying to do without sorting
    # track first max and second max in parallel
    # if a new max is found, update the previous max to the second max, set new max
    # if a new second max is found, update the second max
    # can be done in one iteration. so, O(n)

    if len(lst) < 2:
        raise ValueError("Need at least two values")
    
    # initialize the first and second to the lowest
    first_max, second_max = float('-inf'), float('inf')

    for x in lst:
        # if there are repetitions in the max values,
        # we cannot take the repeated value as second max
        # so check for equality
        if x > first_max:
            # found a new max
            # update first max and reset second max to first max
            second_max, first_max = first_max, x
        elif x > second_max and x != first_max:
            second_max = x
    
    if second_max == float('-inf'):
        raise ValueError("Invalid Input. Found only one maximum: {}".format(first_max))
    
    return second_max

def test_find_second_maximum():
    testcases = [
        [[9,2,3,6], 6],
        [[2, 3, 4, 4], 3],
        [[2, 3, 2], 2]
    ]

    for testcase in testcases:
        inp, expected_output = testcase
        actual_output = find_second_maximum(inp)
        assert expected_output == actual_output

test_find_second_maximum()

    
    
