from math import exp


def longest_substring_with_k_distinct_chars(string, K):
    """
    Given a string, find the length of the longest substring in it with no more than K distinct characters.

    Example 1:

    Input: String="araaci", K=2
    Output: 4
    Explanation: The longest substring with no more than '2' distinct characters is "araa".
    Example 2:

    Input: String="araaci", K=1
    Output: 2
    Explanation: The longest substring with no more than '1' distinct characters is "aa".
    Example 3:

    Input: String="cbbebi", K=3
    Output: 5
    Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".
    Example 4:

    Input: String="cbbebi", K=10
    Output: 6
    Explanation: The longest substring with no more than '10' distinct characters is "cbbebi".
    """

    # for each substring
    #   seen := set()
    #   for c in substring:
    #       seen.add(c)
    #       if len(seen) <= K:
    #           longest = max(longest, len(substring))

    # handle invalid inputs
    # K >= 1
    # len(string) >= 1 ?

    longest = 0
    for i in range(len(string)):
        seen = set()
        for j in range(i, len(string)):
            seen.add(string[j])
            if len(seen) <= K:
                longest = max(longest, j-i+1)
    
    return longest

def longest_substring_with_k_distinct_chars_optimized(string, K):
    # the previous solution runs in O(n^2)
    # we find the frequency of substrings multiple times
    # can we do better than O(n^2)?

    # keep a sliding window
    # when to extend the window?
    # a
    #   ar
    #      ara
    #       araa
    #         c -> #distinct chars > K now. c -> is the new end.
    #         what is the new start?
    #           move start from the left.
    #               drop char at each move until #distinct_chars <= k
    #               if #dis
    # when to shrink the window? -> after reaching K, shrink until #distinct < K
    
    window_start = 0
    longest = 0
    char_count = dict()

    for window_end in range(len(string)):
        char_count[string[window_end]] = char_count.get(string[window_end], 0) + 1
        if len(char_count) <= K:
            # update longest length
            longest = max(longest, window_end - window_start + 1)
        else:
            # exceeded beyond K. shrink the window
            # window start can't go beyond the end.
            while len(char_count) > K:
                char_count[string[window_start]] -= 1
                if char_count[string[window_start]] == 0:
                    del char_count[string[window_start]]
                window_start += 1
            
    return longest

def test_longest_substring_with_k_distinct_chars():
    testcases = [
        [["araaci", 2], 4],
        [["araaci", 1], 2],
        [["cbbebi", 3], 5],
        [["cbbebi", 10], 6]
    ]

    for testcase in testcases:
        string, K = testcase[0]
        expected_output = testcase[1]
        actual_output = longest_substring_with_k_distinct_chars_optimized(string, K)
        assert actual_output == expected_output
    
test_longest_substring_with_k_distinct_chars()