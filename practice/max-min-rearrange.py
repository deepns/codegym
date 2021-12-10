def max_min_rearrange(lst):
    # Implement a function called max_min(lst) which will re-arrange
    # the elements of a sorted list such that the 0th index will have
    # the largest number, the 1st index will have the smallest, and
    # the 2nd index will have second-largest, and so on. In other words,
    # all the even-numbered indices will have the largest numbers in the
    # list in descending order and the odd-numbered indices will have the
    # smallest numbers in ascending order.

    # sample input: [1, 2, 3, 4, 5, 6, 7]
    # output: [7, 1, 6, 2, 5, 3, 4]

    # notes:
    #  input list is sorted
    #  output ->. even indices with numbers in descending
    #         ->. odd indices with numbers in ascending
    # 
    # can create a new list
    # with a two pointers tracking the min and max in the first
    # list, append each end to the new list
    #  [1, 2, 3, 4, 5, 6, 7] ->
    #       [7, 1, 2, 3, 4, 5, 6] -> this rearranges by moving elements to the right
    #       [7, 1, 6, 2, 3, 4, 5]
    #       [7, 1, 6, 2, 5, 3, 4]
    #       [7, 1, 6, 2, 5, 4, 3]

    rearranged_lst = []
    min_index, max_index = 0, len(lst) - 1

    while min_index < max_index:
        rearranged_lst.append(lst[max_index])
        rearranged_lst.append(lst[min_index])
        min_index += 1
        max_index -= 1
    
    # if the list length is odd number, we will still
    # have a mid element to take care.
    if len(lst) % 2 == 1:
        rearranged_lst.append(lst[min_index])

    return rearranged_lst

    # do it in place?
    #   TODO

