def find_average_of_contiguous_subarrays(arr, k):
    """
    Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
    """
    # e.g. input
    #   arr: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
    #   return: [2.2, 2.8, 2.4, 3.6, 2.8]

    # array size? 1 <= k < len(arr)
    # brute force
    #   for each window of k in arr:
    #       find the average for window[i..k]

    # check for limits
    #   1 <= k < len(arr)
    #   empty array? => return empty list

    # this can be O(N * k)
    #   if k = N-1, this can grow to O(N^2) in the worst case
    # the internal sum is repeated for each subarray. we can reuse that instead.
    # averages = []
    # for i in range(len(arr) - k + 1):
    #     # sum here iterates the internal array again
    #     averages.append(sum(arr[i:i+k]) / k)   
    # return averages

    if not arr:
        return []

    # averages = []
    # sum_of_window = sum(arr[:k])

    # averages.append(sum_of_window/k)
    # for i in range(1, len(arr) - k + 1):
    #     sum_of_window = sum_of_window - arr[i-1] + arr[i + k - 1]
    #     averages.append(sum_of_window/k)
        
    # return averages

    # still doing the iteration twice. can do better.
    averages = []
    sum_of_window = 0
    start = 0
    for end in range(len(arr)):
        sum_of_window += arr[end]
        if end >= k - 1:
            averages.append(sum_of_window/k)
            sum_of_window -= arr[start]
            start += 1
    
    return averages

def test_find_average_of_contiguous_subarrays():
    # some simple cases
    testcases = [
        [
            [[1, 3, 2, 6, -1, 4, 1, 8, 2], 5],
            [2.2, 2.8, 2.4, 3.6, 2.8]
        ],
        [
            [[15, 11, 9, -5, -8, 8, -1, 1, 5, -2], 4],
            [7.5, 1.75, 1.0, -1.5, 0.0, 3.25, 0.75]
        ]
    ]

    for testcase in testcases:
        arr, k = testcase[0]
        expected_output = testcase[1]
        actual_output = find_average_of_contiguous_subarrays(arr, k)

        assert actual_output == expected_output, f'{actual_output} didn\'t match with {expected_output}'

test_find_average_of_contiguous_subarrays()
        