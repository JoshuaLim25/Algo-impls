package two_pointers

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
// Given integer array nums, sorted in non-decreasing order.
// remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements SHOULD BE KEPT THE SAME.
// Then return the number of UNIQUE ELEMENTS in nums.

// BRUTE:
// Decisions:
// @ element i, check if elem in set. if so, remove it from array and decrement numUnique, and if not, add it to the set and incr numUnique. Continue on.

// NOTE: ordering matters.
// INFO: think set, so map[int]bool
// GIVEN: [ 2 2 3 4 4 5 ] -> 2 (3 and 5), and [ 2 3 4 5 ]

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
//     Return k.

import "fmt"
