from collections import deque

def find_product(lst):
    # Implement a function, find_product(lst), which modifies a list
    # so that each index has a product of all the numbers present in
    # the list except the number stored at that index.

    # e.g. arr = [1,2,3,4]
    # result = [24, 12, 8, 6]

    # notes:

    # inputs can be integers or floating point numbers
    # can there be zeros? if one of the number is zero, then all
    # values in the result would be zero except the one where we
    # have zero.
    
    # brute force
    # for each index i in lst
    #   for each index j in lst
    #       prod *= lst[j] if i != j
    #   lst[i] = prod

    # this will be O(n^2). can we do better?

    # clearly we're doing repeated work in the inner loop
    # that can be leveraged.
    # product_of_all_numbers_except_at_i = product_of_all_numbers_before_i * product_of_all_numbers_after_i

    # the above is true for all index i in lst
    # if we maintain two lists, one that tracks product before i and other after i
    # then result can be calculated from those two lists.

    # product_of_all_numbers_before_i and product_of_all_numbers_after_i
    # can be calculated linearly. so, the result can be calculated in
    # linear time too instead of quadratic time.

    # some edge cases to take care
    # empty list?
    # single element
    # zeros?

    # e.g.
    # [2, 5, 7, -1]

    # product_of_all_before_i = [1, 2, 10, 70]
    # product_of_all_after_i = [-35,-7,-1,1]

    product_of_all_numbers_before_i = [1] # this needs to be calculated from 1..(n-1)
    product_of_all_numbers_after_i = deque([1]) # this needs to be calculated from 0...n-2
                                                # can initialize another array of same length too
                                                # using deque makes the iteration easier and avoids
                                                # additional walks for the initiatialization.

    for x in lst[:-1]:
        product_of_all_numbers_before_i.append(x * product_of_all_numbers_before_i[-1])
    
    # have to walk from the end
    for x in reversed(lst[1:]):
        product_of_all_numbers_after_i.appendleft(x * product_of_all_numbers_after_i[0])

    # don't need to create a new list to carry the result
    # can use one of the two product lists created.
    for i in range(len(product_of_all_numbers_before_i)):
        product_of_all_numbers_before_i[i] *= product_of_all_numbers_after_i[i]

    return product_of_all_numbers_before_i

def test_find_product():
    testcases = [
        [[2, 5, 7, -1], [-35, -14, -10, 70]],
        [[1,2,3,4], [24, 12, 8, 6]],
        [[2, 5, 9, 3, 6], [810, 324, 180, 540, 270]],
        [[0, 1, 10, 100], [1000, 0, 0, 0]],
        [[0, 2, 9, 0, 12, 25], [0, 0, 0, 0, 0, 0]]
    ]

    for testcase in testcases:
        inp, expected_output = testcase
        actual_output = find_product(inp)
        assert actual_output == expected_output, "Inp: {}, Exp: {}, Actual: {}".format(inp, expected_output, actual_output)

test_find_product()
