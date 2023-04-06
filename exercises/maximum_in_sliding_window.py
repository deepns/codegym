# Given an integer list nums, find the maximum values in all the
# contiguous subarrays (windows) of size w.
# 
# Constraints:
#   array length: 1..10^3
#   array values: -10^4...10^4
#   window size: 1..len(arr)
#
# e.g. 
#   -4 2 -5 3 6
#   output: 2, 3, 6
#
#   -1 -1 -1 -1 -1
#   -1 -1

# for each window..compute maximum value (iterate through the window)
# n-k window and k iterations in each window.. so (n-k)*k in time
# can the k iterations within the window be avoided?
#
# at the very least, compute max for the first window
# when the window is slided, need to find max for that window
# when will that change? if new entry is > max_seen_so_far
# what if the max_seen_so_far goes out of the window
#
#   -4 2 -5 3 1 4 2, 2 1, k = 3
#   2, 3, 3, 4, 4, 4, 2, 
# if i < k:
#   compute max_seen_so_far
# else:
#   max_seen_so_far = max(max_seen_so_far, new_entry)
# Checkout https://leetcode.com/problems/sliding-window-maximum/

def find_max_sliding_window_brute_force(nums:str, w:int):
    sliding_window_maximums = []
    n = len(nums)
    for i in range(n - w):
        sliding_window_maximums.append(max(nums[i:i+w]))

    return sliding_window_maximums

# TODO
# add an optimized version of the solution