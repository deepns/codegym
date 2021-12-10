def merge_lists(lst1, lst2):
    # Write your code here
    # merge two sorted lists
    #
    # lst1 = [1, 3, 4, 6]
    # lst2 = [-4, -2, 4, 6, 7]

    # return [-4, -2, 1, 3, 4, 4, 6, 6, 7]

    # both lists are already sorted
    
    # constraints? how big the list?
    # each list have to be walked fully. O(m) + O(n) => O(m + n)

    # one of the list can be empty too
    # can return the non-empty list

    # walk two list in parallel.
    # pick the smaller one in each iteration

    # some optimizations
    if not lst1:
        return lst2
    
    if not lst2:
        return lst1

    mergedlist = []
    i, j = 0, 0
    # length of traversal is limited by the shorter of the two lists
    # if the list were to be updated in place, then each insertion will
    # require walking down the list to be inserted. that would make
    # complexity to be O((m+n)*m) if updating list1 in place and list2 has all elements smaller than list2
    while (i < len(lst1) and j < len(lst2)):
        if lst1[i] <= lst2[j]:
            mergedlist.append(lst1[i])
            i += 1
        else:
            mergedlist.append(lst2[j])
            j += 1

    # there can be left over items if the lists are not equal length
    
    # may be not so pythonic. can add the list instead.
    # while i < len(lst1):
    #     mergedlist.append(lst1[i])
    #     i += 1
    
    # while j < len(lst2):
    #     mergedlist.append(lst2[j])
    #     j += 1
    mergedlist.extend(lst1[i:]) # same as mergedlist += lst1[i:]
    mergedlist.extend(lst2[j:]) # same as mergedlist += lst2[j:]
    
    return mergedlist

def test_merged_list():
    import random
    NUM_TESTS = 10
    for i in range(NUM_TESTS):
        lst1size = random.randint(0, 10)
        lst2size = random.randint(0, 10)

        lst1 = sorted([random.randint(-100, 100) for _ in range(lst1size)])
        lst2 = sorted([random.randint(-100, 100) for _ in range(lst2size)])

        print("Merging", lst1, lst2)
        mergedlist = merge_lists(lst1, lst2)

        assert mergedlist == sorted(lst1 + lst2), "Merged list is not sorted"

test_merged_list()