def rearrange(lst):
    # Implement a function rearrange(lst) which rearranges the
    # elements such that all the negative elements appear on the
    # left and positive elements appear at the right of the list.
    # Note that it is not necessary to maintain the sorted order
    # of the input list.

    # sample input: [10,-1,20,4,5,-9,-6]
    # output: [-1,-9,-6,10,20,4,5]

    # Zero? - treat it as positive
    # rearrange in place? or create new list?

    # sort? nlogn
    #   naturally separates positive and negative numbers
    # can we do better?

    # creating new list is fairly straightforward
    # move positive numbers to one list and negative numbers to another
    # combine the two lists, can be done in O(n)

    # positives = [x for x in lst if x >= 0]
    # negatives = [x for x in lst if x < 0]

    # return negatives + positives

    # rearrange in place?
    #   walk the list
    #   if lst[i].ispositive()
    #       swap or move?.
    # [10,-1,20,4,5,-9,-6]
    #   any removal or insertion other than the end will require
    #   reordering of the other elements as well.
    #   swapping may work.
    #   to swap, we need to track two pos. left_most_positive
    #   (if moving negative to the left)
    #   or right_most_negative (if moving positive to the right)

    # left_most_positive_index = 0
    # [-1, 10, 20, 4, 5, -9, -6] # swap(0, 1), left_most_positive_index=1
    # [-1, -9, 20, 4, 5, 10, -6] # swap(1, 5), left_most_positive_index=2
    # [-1, -9, -6, 20, 4, 5, 10]

    # [-1, -2, 4, 5, -4]
    #   [-1, -2, ]

    left_most_positive = 0
    for i in range(len(lst)):
        if lst[i] < 0:
            if i != left_most_positive:        
                # swap if a positive has been seen and i > positive index
                lst[i], lst[left_most_positive] = lst[left_most_positive], lst[i]
            left_most_positive += 1
    
    return lst

    



    
