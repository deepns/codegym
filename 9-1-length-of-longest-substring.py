def length_of_longest_substring(string, k):
    """
    Given a string with lowercase letters only, if you are allowed to replace no more than k letters with any letter, find the length of the longest substring having the same letters after replacement.

    Example 1:

    Input: String="aabccbb", k=2
    Output: 5
    Explanation: Replace the two 'c' with 'b' to have the longest repeating substring "bbbbb".
    Example 2:

    Input: String="abbcb", k=1
    Output: 4
    Explanation: Replace the 'c' with 'b' to have the longest repeating substring "bbbb".
    Example 3:

    Input: String="abccde", k=1
    Output: 3
    Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".
    """
    # for each substring
    #   if char seen is same as prev or first char:
    #       update max length
    #       continue
    #   else
    #       if substitutions < k:
    #           substitute current char with prev char
    #           increment substitutions
    #           update max length
    #       else:
    #           can't do any more substitutions. stop searching this substring

    # e.g. aabccbb, k=2
    #   aa -> ls=2
    #   aab -> substitutions=1, ls=3
    #   aabc -> substitutions=2, ls=4
    #   aaaac -> can't do any more substitutions
    #
    #   substring starting at b
    # ababcbdb, k=2, output => 7

    max_length = 0
    for i in range(len(string)):
        substitutions = 0
        for j in range(i, len(string)):
            # process the substring
            if string[j] == string[i]:
                max_length = max(max_length, j-i+1)
            else:
                if substitutions < k:
                    substitutions += 1
                    max_length = max(max_length, j-i+1)
                else:
                    break
    return max_length

def test_length_of_longest_substring():
    testcases = [
        [["aabccbb", 2], 5],
        [["abbcb", 1], 4],
        [["abccde", 1], 3],
        [["ababcbbb", 2], 7]
    ]
    
    for testcase in testcases:
        string, k = testcase[0]
        expected_output = testcase[1]

        actual_output = length_of_longest_substring(string, k)
        assert actual_output == expected_output, \
            f'inp:({string}, {k}), exp: {expected_output}, actual:{actual_output}'

test_length_of_longest_substring()