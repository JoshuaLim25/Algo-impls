/*
55. Jump Game - Medium

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

* NOTE: it's NOT additive. i.e., for [2,3,1,1,4], you can jump to idx 1 or 2, but if you jump to idx 1 your maxJump isn't 3+1. It's just the number @ that idx.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

Constraints:
    1 <= nums.length <= 104
    0 <= nums[i] <= 105
*/

package com.algos.greedy;

public class JumpGame55 {
    // basically, try to just take a SINGLE step.
    // if you encounter a better option, switch to that option
    // otherwise, keep carrying on
    // you "make it" to the end if you don't totally run out before then (-1)
    // Consider when i == n-2 and fuel == 0. When i becomes n-1, fuel will be -1
    public static boolean canJump(int[] nums) {
        int fuel = 0;
        for (int num : nums) {
            if (fuel < 0) return false;
            else if (num > fuel) fuel = num;
            --fuel;
        }
        return true;
    }

    public static void main(String[] args) {
        int[] nums = {2,3,1,1,4};
        System.out.println(String.format("true == %b", JumpGame55.canJump(nums)));

        int[] nums2 = {3,2,1,0,4};
        System.out.println(String.format("false == %b", JumpGame55.canJump(nums2)));

        int[] nums3 = {2,5,0,0};
        System.out.println(String.format("true == %b", JumpGame55.canJump(nums3)));

    }
}
