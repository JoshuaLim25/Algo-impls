/*
347. Top K Frequent Elements - Medium

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]
*/

import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Queue;

public class Top_K_Frequent_Elements347 {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> freqMap = new HashMap<>();
        // idea: init a minHeap. Add frequencies encountered for n in nums
        // add all freqs from map to the heap
        for (int n : nums) {
            // if (!freqMap.containsKey(n)) {
            // freqMap.put(n, 1); // BAD: not 0
            // } else {
            // int freq = freqMap.get(n);
            // freqMap.put(n, ++freq);
            // }
            freqMap.put(n, freqMap.getOrDefault(n, 0) + 1);
        }

        // INFO: to get CUSTOM ORDERING based on frequency:
        Queue<Integer> minHeap = new PriorityQueue<>(
                (a, b) -> freqMap.get(a) - freqMap.get(b));

        for (int n : freqMap.keySet()) {
            minHeap.add(n);
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }

        int[] res = new int[k];
        for (int i = k - 1; i >= 0; --i) {
            res[i] = minHeap.poll();
        }
        return res;
    }

    public static void main(String[] args) {
        Top_K_Frequent_Elements347 sol = new Top_K_Frequent_Elements347();
        int[] nums = { 1, 1, 1, 2, 2, 3 };
        int k = 2;
        System.out.println(String.format("[1, 2] == %s", Arrays.toString(sol.topKFrequent(nums, k))));

        int[] nums2 = { 1 };
        int k2 = 1;
        System.out.println(String.format("[1] == %s", Arrays.toString(sol.topKFrequent(nums2, k2))));
    }
}
