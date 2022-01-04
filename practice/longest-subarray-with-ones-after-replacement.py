def length_of_longest_substring(arr, k):
    """
    Given an array containing 0s and 1s, if you are allowed to replace no more than 'k'
    0s with 1s, find the length of the longest contiguous subarray having all 1s.

    Example 1:

    Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
    Output: 6
    Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.
    Example 2:

    Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
    Output: 9
    Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
    """

    # replace 0 with 1, up to k times
    # find the length of the substring with contiguous 1s
    # [0, 1, 0, 0, 1, 1, 0, 0, 0], k = 2. if we replace 0s at index 2 and 3,
    # length of the substring(1, 5) will be 5. 

    # for each substring.. we can do k replacements. the replacement can
    # be anywhere.. we have to try all positions. in substring [0, 1, 0, 0, 0, 1]
    # we can replace [1, 1, 1, 0, 0, 1] or [0, 1, 1, 1, 0, 1]

    # brute forcing with each substring.
    # will be O(N^2)

    len_of_longest_ones = 0

    N = len(arr)
    for i in range(N):
        replacements = 0
        for j in range(i, N):
            if arr[j] == 0 and replacements >= k:
                # quit this substring if we crossed the replacement limit
                break

            # XXX This is flawed. the replacement can be done further
            # down the substring too. this solution would fail.
            if arr[j] == 0:
                replacements += 1

            len_of_longest_ones = max(len_of_longest_ones, j - i + 1)
    
    return len_of_longest_ones

def length_of_longest_substring_v1(arr, k):
    # using a sliding window method?
    # window tracks the start and end of substring with 1s
    # extend the window if arr[i] == 1 or arr[i] == 0 and replacements < k

    # some corner cases, arr with all ones or 0s.
    
    N = len(arr)
    window_start = 0
    len_of_longest_ones = 0
    num_ones_in_substr = 0
    
    for window_end in range(N):
        if arr[window_end] == 1:
            # extend the window without replacement
            num_ones_in_substr += 1
        
        window_length = window_end - window_start + 1
        if (window_length - num_ones_in_substr) > k:
            # we've exceeded the replacements. time to shrink the window
            if arr[window_start] == 1:
                num_ones_in_substr -= 1    
            window_start += 1
            window_length = window_end - window_start + 1
        
        len_of_longest_ones = max(len_of_longest_ones, window_length)
    
    return len_of_longest_ones            

def test_length_of_longest_substring():
    testcases = [
        [
            [[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], 2], 6
        ],
        [
            [[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], 3], 9
        ],
        [
            [[0, 1, 1, 0, 0], 2], 4
        ],
        [
            [[0, 0, 0, 0, 0], 3], 3
        ],
        [
            [[0, 0, 0, 0, 0], 0], 0
        ],
        [
            [[1, 1, 1, 1, 1], 3], 5
        ]
    ]

    for testcase in testcases:
        arr, k = testcase[0]
        expected_output = testcase[1]

        actual_output = length_of_longest_substring_v1(arr, k)
        assert expected_output == actual_output, f'actual: {actual_output}, expected: {expected_output}, inp: arr={arr}, k={k}'

test_length_of_longest_substring() 
