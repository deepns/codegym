def find_sum(lst, k):
    # lst = [1,21,3,14,5,60,7,6]
    # k = 81

    # return [21, 60]

    # naive.
    # for each number x in lst
    #    check k-x is in lst and index(k-x) is not index(x)
    # this will be O(n^2)

    # can we do better?
    #   we are doing many look ups here. perhaps, use a cache to store looked up numbers?

    # for x in lst:
    #   if k-x in seen_list:
    #       return [x, k-x]
    #

    # what if k is negative? and numbers are negative?
    # is space a constraint?

    # can there be multiple combinations?
    # for e..g [1, 4, 5, 7, 3] and k = 8, then [1, 7] and [5, 3] are valid results

    # will there always exist two numbers for a given k?

    # If the list is empty? -> handled at the for loop. will return an empty list.

    # some other ways to solve this.
    # the lookup for (k-x) can be done through binary search if the list 
    # is sorted.

    seen = set()

    for x in lst:
        if (k - x) in seen:
            return [x, k-x]
        else:
            seen.add(x)
    
    return []

def test_find_sum():
    import random
    NUM_TESTS=5
    TEST_LIST_SIZE=10

    for i in range(NUM_TESTS):
        lst = [random.randint(-100, 100) for _ in range(TEST_LIST_SIZE)]

        # pick two different indices in the list
        i = random.randint(0, TEST_LIST_SIZE-1)
        j = i
        while j == i:
            j = random.randint(0, TEST_LIST_SIZE-1)
        
        k = lst[i] + lst[j]

        expected = [lst[i], lst[j]]
        print("Calling find_sum({}, {}), expecting={}".format(lst, k, expected))
        actual = find_sum(lst, k)
        # since the inputs are randomly generated, there can be multiple pairs
        # that sum to k.
        assert (sorted(actual) == sorted(expected) or sum(actual) == k), \
                "Expected: lst({}, {}) = {}".format(expected, k, actual)
        
test_find_sum()