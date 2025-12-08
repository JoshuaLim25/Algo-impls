/*
217. Contains Duplicate - Easy

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true
Explanation: The element 1 occurs at the indices 0 and 3.
 */

package com.algos.strings;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class ContainsDuplicate217 {
    public static boolean contains(List<Integer> nums) {
        Set<Integer> set = new HashSet<>();
        for (int n : nums) {
            if (set.contains(n)) return true;
            set.add(n);
        }
        return false;
    }

    public static void main(String[] args) {
// [1,2,3,1]
        List<Integer> nums = new ArrayList<>(List.of(1,2,3,1));
        System.out.println(ContainsDuplicate217.contains(nums));
    }
}
