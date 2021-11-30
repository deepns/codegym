def maximum_sum_subarray(arr, k):
    """
    Given an array of positive numbers and a positive number ‘k,’ find the maximum sum of any contiguous subarray of size ‘k’.

    Example 1:

    Input: [2, 1, 5, 1, 3, 2], k=3 
    Output: 9
    Explanation: Subarray with maximum sum is [5, 1, 3].
    Example 2:

    Input: [2, 3, 4, 1, 5], k=2 
    Output: 7
    Explanation: Subarray with maximum sum is [3, 4].
    """

    # arr: array of positive numbers
    # 0 < k < len(arr)

    # for each subarray of size k in arr:
    #   find sum of subarray.
    #   update max

    # this will be O(N^2) in time
    # can do it in O(N) by using sliding window

    max_sum = float('-inf')
    sub_array_sum = 0
    sub_array_start = 0

    for sub_array_end in range(len(arr)):
        sub_array_sum += arr[sub_array_end]

        # if the window size is reached, slide to the right
        # update max seen so far
        if sub_array_end >= k - 1:
            max_sum = max(max_sum, sub_array_sum)
            sub_array_sum -= arr[sub_array_start]

            # slide to the right by moving the starting point
            sub_array_start += 1
    
    return max_sum

def test_maximum_sum_subarray():
    testcases = [
        [
            [[2, 1, 5, 1, 3, 2], 3], 9
        ],
        [
            [[2, 3, 4, 1, 5], 2], 7
        ],
        [
            [[2, 4, 5, 7, 11, 3], 1], 11
        ]
    ]
    for testcase in testcases:
        arr, k = testcase[0]
        expected_output = testcase[1]
        actual_output = maximum_sum_subarray(arr, k)

        assert actual_output == expected_output, f'{actual_output} didn\'t match with {expected_output}'

test_maximum_sum_subarray()