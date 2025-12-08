/*
167. Two Sum II - Input Array Is Sorted - Medium

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

*You may not use the same element twice.*              => while(l < r) { ... }
*Your solution must use only constant extra space.*    => two pointers
*The tests are generated such that there is exactly one solution.*  => no dupes 
 */

package com.algos.twoPointers;

import java.util.ArrayList;
import java.util.List;

record PairInt(int first, int second) {
}

public class TwoSumSorted167 {
    // just remember to idx++ when returning since 1-indexed
    public static PairInt twoSum(List<Integer> nums, int targetSum) {
        if (nums.size() <= 1)
            return new PairInt(1, 1);
        // input: nums = [ 2 7 11 15 ], targetSum = 9;
        // ↑ ↑
        // 1. nums[l] + nums[r];
        // 2. Check if sum < targetSum. If so l++;
        // 3. Check if sum > targetSum. If so r--;
        // 4. Check if sum == targetSum. If so return l, r;
        int l = 0, r = nums.size() - 1;

        while (l < r) { // "You may not use the same element twice."
            int sum = nums.get(l) + nums.get(r);
            System.err.println("DEBUG[1]: TwoSumSorted167.java:34: sum=" + sum);
            if (sum == targetSum) {
                return new PairInt(l+1, r+1);
            } else if (sum < targetSum)
                l++;
            else if (sum > targetSum)
                r--;
        }
        return new PairInt(1, 1);
    }

    public static void main(String[] args) {
        var nums1 = new ArrayList<>(List.of(2, 7, 11, 15));
        int target1 = 9;
        System.out.println(String.format("[1, 2] == %s", TwoSumSorted167.twoSum(nums1, target1)));

        var nums2 = new ArrayList<>(List.of(2, 3, 4));
        int target2 = 6;
        System.out.println(String.format("[1, 3] == %s", TwoSumSorted167.twoSum(nums2, target2)));
    }
}
