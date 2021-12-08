import collections


def longest_substring_with_distinct_chars(string):
    """
    Given a string, find the length of the longest substring, which has all distinct characters.

    Example 1:

    Input: String="aabccbb"
    Output: 3
    Explanation: The longest substring with distinct characters is "abc".
    Example 2:

    Input: String="abbbb"
    Output: 2
    Explanation: The longest substring with distinct characters is "ab".
    Example 3:

    Input: String="abccde"
    Output: 3
    Explanation: Longest substrings with distinct characters are "abc" & "cde".
    """
    # in each substring
    #   for each char in substring
    #       if char is distinct # char is not seen yet
    #           longest_substr = max(longest_substr, len_of_substr)
    #       else
    #           break
    # this runs in O(N^2), space O(1). can get very slow for significantly large N

    # handle edge cases 
    # TODO
    #   empty string? => return 0
    #   all chars are same (e.g. "aaaaaaaaaaaaaaa")
    #   all chars are different (e.g. "abcdedghigkl")
    #   charset -> only lower case?

    len_of_longest_substr_with_distinct_chars = 0
    for i in range(len(string)):
        seen = set()
        for j, char in enumerate(string[i:], start=i):
            if char not in seen:
                # unique substr so far. add char to the seen.
                seen.add(char) # size of this set is limited by the charset. 
                len_of_longest_substr_with_distinct_chars = max(len_of_longest_substr_with_distinct_chars, j - i + 1)
            else:
                # found duplicate char. move to next substr starting at i+1
                break    
    return len_of_longest_substr_with_distinct_chars

def longest_substring_with_distinct_chars_optimized(string):
    # the last solution ran in O(N^2)
    # can definitely do better than that.
    # substr -> sliding window
    #   keep extending the window until a char is repeated
    #   if a char is repeated, all the substr thus far will become invalid.
    #   update the length.
    #   restart the window from the repeated char.

    seen = collections.Counter()
    length_of_longest_substr = 0
    window_start = 0
    for window_end, char in enumerate(string, start=0):
        seen[char] += 1
        if seen[char] > 1:
            # found a duplicate.
            # clear the seen chars so far and restart from the current char.
            # shrink the window until the current char is not repeated
            while seen[char] > 1:
                seen[string[window_start]] -= 1
                window_start += 1
        else:
            length_of_longest_substr = max(length_of_longest_substr, window_end - window_start + 1)

    return length_of_longest_substr

def test_longest_substring_with_distinct_chars():
    testcases = [
        ["aabccbb", 3],
        ["abbbbb", 2],
        ["abccde", 3],
        ["aaaaaaaaaaaaaaaaaaaa", 1],
        ["abcdefghijklmnopqrstuvwxyz", 26],
        ["pwwkew", 3],
        ["dvdf", 3],
        ["davdf", 4]
    ]

    for testcase in testcases:
        string, expected_output = testcase
        actual_output = longest_substring_with_distinct_chars(string)
        assert expected_output == actual_output, f'Inp: {string}, exp: {expected_output}, actual: {actual_output}'

        actual_output = longest_substring_with_distinct_chars_optimized(string)
        assert expected_output == actual_output, f'Inp: {string}, exp: {expected_output}, actual: {actual_output}'

test_longest_substring_with_distinct_chars()