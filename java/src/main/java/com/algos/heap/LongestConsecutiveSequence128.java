package com.algos.heap;

/*
128. Longest Consecutive Sequence - Medium

Given an **unsorted** array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
- I.e., no sorting allowed

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Example 3:
Input: nums = [1,0,1,2]
Output: 3
*/

import java.util.PriorityQueue;
import java.util.Queue;

public class LongestConsecutiveSequence128 {
    public static int longestConsecutive(int[] nums) {
        int n = nums.length;   // NOTE: NEED SINCE RANGING MODIFIES HEAP LEN
        if (nums == null || n == 0) return 0;
        int total = 1, curLen = 1;
        Queue<Integer> minHeap = new PriorityQueue<>();

        // [100,4,200,1,3,2]
        // 1 2 3 4 
        //       i             
        //   curLen = 3, total = 3
        for (int num : nums) {
            minHeap.add(num);
        }


        int last = minHeap.poll();
        // BAD: for (int i = 1; i < minHeap.size(); ++i) {
        for (int i = 1; i < n; ++i) {
            int cur = minHeap.poll();
            if (cur == last + 1) {
                ++curLen;
                total = Math.max(total, curLen);
            } else if (cur == last) {
                continue;
            } else {
                curLen = 1;
            }
            // NOTE: dont forget to update this
            last = cur;
        }

        return total;
    }

    // public static void main(String[] args) {
    //     int[] nums = {100,4,200,1,3,2};
    //     int soln = Solution.longestConsecutive(nums);
    //     System.out.println(String.format("4 == %d", soln));
    //
    //     int[] nums2 = {0,3,7,2,5,8,4,6,0,1};
    //     int soln2 = Solution.longestConsecutive(nums2);
    //     System.out.println(String.format("9 == %d", soln2));
    //
    //
    //     int[] nums3 = {1,0,1,2};
    //     int soln3 = Solution.longestConsecutive(nums3);
    //     System.out.println(String.format("3 == %d", soln3));
    // }
}
