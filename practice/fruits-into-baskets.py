import collections


def fruits_into_baskets(fruits):
    """
    Given an array of characters where each character represents a fruit tree, you are given two baskets,
    and your goal is to put maximum number of fruits in each basket. The only restriction is that each
    basket can have only one type of fruit.

    You can start with any tree, but you can’t skip a tree once you have started. You will pick one fruit
    from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.

    Write a function to return the maximum number of fruits in both baskets.

    Example 1:

    Input: Fruit=['A', 'B', 'C', 'A', 'C']
    Output: 3
    Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']
    Example 2:

    Input: Fruit=['A', 'B', 'C', 'B', 'B', 'C']
    Output: 5
    Explanation: We can put 3 'B' in one basket and two 'C' in the other basket. 
    This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']
    """

    # we could count them and return the top two. but that will violate the
    # restriction of starting from one point and not picking the subsequent ones

    # for each subarray in the fruits:
    #   start from fruit_i
    #       if fruit_i not in basket:
    # #         if len(basket) < 2:
    #               # good to go
    #               add fruit_i into the basket
    #           else
    #               break
    #      else:
    #           add fruit into the basket
    #           update max fruits

    #       if len(basket) < 2 or fruit_i in basket:
    #           add fruit_i into basket
    #           update max count 
    #       else:
    #           break

    # take care of edge cases here
    # fruits must not be empty

    max_fruits = 0
    MAX_FRUIT_TYPES_IN_BASKET = 2
    for i in range(len(fruits)):
        basket = set()
        for j in range(i, len(fruits)):
            # fruit case sensitive? ignore it for now.
            if len(basket) < MAX_FRUIT_TYPES_IN_BASKET or fruits[j] in basket:
                # we could also use a dict instead of set, and count the number
                # of each fruit in each iteration
                basket.add(fruits[j])
                max_fruits = max(max_fruits, j - i + 1)
            else:
                # exceeded the max fruit types that can go in our buckets for this subarray
                break
    
    return max_fruits

def fruits_into_baskets_optimized(fruits):
    # the above solution runs in O(N^2). space -> O(1), since we put only two types of
    # fruits in our basket
    # can we do it less than that?
    # suppose we keep a sliding window within the fruits array
    #   the window keeps track of the fruits that can go in our basket
    #   add a fruit always
    #   if len(basket) exceeds 2:
    #       shrink the window by dropping fruits from the start
    # wait there could be a catch here
    #   A, B, A, B, C
    #   when we get to C, 
    #       basket = {A, B, C}
    #       start = 0
    #       remove fruits[start] from basket. 
    #       basket can't be set here. we will lose the count of fruits
    #       it has to be dict. move start until basket[fruit] goes to zero

    max_fruits = 0
    basket = collections.Counter()
    window_start = 0

    for window_end in range(len(fruits)):
        fruit = fruits[window_end]
        basket[fruit] += 1

        if len(basket) <= 2:
            # update max
            max_fruits = max(max_fruits, window_end - window_start + 1)
        else:
            # shrink the window
            while len(basket) > 2:
                fruit = fruits[window_start]
                basket[fruit] -= 1
                if basket[fruit] == 0:
                    del basket[fruit]
                window_start += 1
    
    return max_fruits


def test_fruits_into_baskets():
    testcases = [
        [['A', 'B', 'C', 'B', 'B', 'C'], 5],
        [['A', 'B', 'C', 'A', 'C'], 3],
        [['A', 'B', 'A', 'B', 'C', 'A', 'C'], 4]
    ]

    for testcase in testcases:
        fruits, expected_max_fruits = testcase[0], testcase[1]
        actual_max_fruits = fruits_into_baskets(fruits)
        assert actual_max_fruits == expected_max_fruits

        actual_max_fruits = fruits_into_baskets_optimized(fruits)
        assert actual_max_fruits == expected_max_fruits

test_fruits_into_baskets()