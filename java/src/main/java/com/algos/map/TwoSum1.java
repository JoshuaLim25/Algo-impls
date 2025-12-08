/*
1. Two Sum - Easy
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 */

import java.util.HashMap;
import java.util.List;
import java.util.Map;

record Pair(int first, int second) {
}

public class TwoSum1 {
    public static Pair twoSum(List<Integer> nums, int target) {
        Map<Integer, Integer> diffs = new HashMap<>();
        for (int i = 0; i < nums.size(); ++i) {
            int diff = target - nums.get(i);
            if (diffs.containsKey(diff)) {
                return new Pair(i, diffs.get(diff));
            }
            diffs.put(nums.get(i), i); // mapping from num to idx
        }
        return new Pair(0, 0);
    }

    public static void main(String[] args) {
        var nums = List.of(2, 7, 11, 15);
        var target = 9;
        System.out.println(TwoSum1.twoSum(nums, target));
    }
}
