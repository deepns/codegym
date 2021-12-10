def remove_even(lst):
    # Given a list of numbers, remove all even numbers from the list
    # Given a list my_list = [1,2,4,5,10,6,3]
    # return my_list = [1,5,3]

    # remove elements from the given list. do not create a new list.
    # creating a new list is more efficient, if space is not a constraint.
    # tradeoff between time and space.

    # updating existing list requires lot of move
    #   should the order be preserved?
    #   what is the cost of clearing the list?
    # inputs:
    #   list may contain all odd numbers - no change to the list
    #   may be all even - remove all elements from the list
    #   may be equally split between odd and even

    # basic
    #   filter the list by including only odd numbers
    # return list(filter(lambda x: x % 2 == 1, lst))

    # updating the list in place requires lot of move
    # update the list while iterating can create undeterministic results
    
    # list.remove(value) # removes the first occurrence of the value.
    # that includes searching the list as well.

    # list.pop(i) -> removes element at index i

    # replace the elements in existing list
    # lst[::] = list(filter(lambda x: x % 2 == 1, lst))
    
    # or use list comprehension
    # return [x for x in lst if x % 2 != 0]

    # update the list in place
    i = 0
    while i < len(lst):
        if (lst[i] % 2 == 0):
            lst.pop(i)
        else:
            i += 1

def test_remove_even():   
    import random
    alleven = [random.randint(1, 100) * 2 for _ in range(10)]
    allodds = [(random.randint(1, 100) * 2) + 1 for _ in range(10)]
    oddeven = [random.randint(1, 100) for _ in range(10)]

    for lst in [alleven, allodds, oddeven]:
        print(lst, "==>", end="")
        remove_even(lst)
        print(lst)
        assert all([ x % 2 == 1 for x in lst]), "Found even numbers in the list"