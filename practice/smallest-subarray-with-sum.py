def smallest_subarray(arr, S):
    """
    Given an array of positive numbers and a positive number ‘S,’ 
    find the length of the smallest contiguous subarray whose sum is
    greater than or equal to ‘S’. Return 0 if no such subarray exists.

    Example 1:

    Input: [2, 1, 5, 2, 3, 2], S=7 
    Output: 2
    Explanation: The smallest subarray with a sum greater than or equal to '7' is [5, 2].
    Example 2:

    Input: [2, 1, 5, 2, 8], S=7 
    Output: 1
    Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].
    Example 3:

    Input: [3, 4, 1, 1, 6], S=8 
    Output: 3
    Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1] 
    or [1, 1, 6].
    """

    # arr: positive numbers
    # S: > 0

    # find length of smallest contiguous subarray with sum >= S
    # if no subarray with sum >= S exists, return 0.

    # for each subarray in arr:
    #   find sum
    #   if sum_of_sub_array >= S:
    #       length = sub_array_end - sub_array_start
    #       length_of_smallest_subarray = min(length_of_smallest_subarray, length)

    # generating subarray with two loop
    #   array of len=4
    # [0], [0, 1], [0, 1, 2], [0, 1, 2, 3]
    # [1], [1, 2], [1, 2, 3]
    # [2], [2, 3]
    # [3]

    length_of_smallest_subarray = float('inf')
    for i in range(len(arr)):
        sum_of_this_subarray = 0
        for j in range(i, len(arr)):
            sum_of_this_subarray += arr[j]
            if sum_of_this_subarray >= S:
                length_of_smallest_subarray = min(length_of_smallest_subarray, j - i + 1)
    
    return 0 if length_of_smallest_subarray == float('inf') else length_of_smallest_subarray

def test_smallest_subarray():
    testcases = [
        [
            [[2, 1, 5, 2, 3, 2], 7],
            2
        ],
        [
            [[2, 1, 5, 2, 8], 7],
            1
        ],
        [
            [[3, 4, 1, 1, 6], 8],
            3
        ],
        [
            [[3, 4, 1, 1, 6], 18],
            0
        ],
        [
            [[3, 4, 1, 1, 6, 3], 18],
            6
        ]
    ]

    for testcase in testcases:
        arr, S = testcase[0]
        expected_output = testcase[1]
        actual_output = smallest_subarray(arr, S)
        assert expected_output == actual_output, \
            f'expected={expected_output}, actual={actual_output} for arr={arr}, S={S}'

test_smallest_subarray()

