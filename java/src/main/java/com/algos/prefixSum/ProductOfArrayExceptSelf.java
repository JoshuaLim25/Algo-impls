/*
238. Product of Array Except Self
Solved
Medium
Topics
premium lock iconCompanies
Hint

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

 

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Algo:
1. Make two lists, suffix and prefix.
2. Keep the first or last number, respectively.
3. Start the loop one past/before that elt.
    4. For the prefix loop, you multiply the last item put in prefix (i-1) and multiply by the nums[idx]
Input: nums = [    1     2     3      4    ]
      prefix: [    1     1     2      6    ]  # 1, 1*1, 1*2, 3*2
      suffix: [    24    12    4     1    ]    # 12*2 <-  4*3  <- 1*4 <- 1

 */

package com.algos.prefixSum;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

// NOTE: calculating products, so need to start @ second/second-last elts
class ProductOfArrayExceptSelf {
    public static List<Integer> doIt(List<Integer> nums) {
        int n = nums.size();
        var prefix = new ArrayList<>(Collections.nCopies(n, 0));
        var suffix = Stream.generate(() -> 0).limit(n).collect(Collectors.toList());
        var res = Stream.generate(() -> 0).limit(n).collect(Collectors.toList());

        prefix.set(0, 1);
        suffix.set(n - 1, 1);

        // think of it as "both lag one behind"
        for (int i = 1; i < n; ++i) {
            prefix.set(i, nums.get(i-1) * prefix.get(i-1));
        }

        for (int i = n - 2; i >= 0; --i) {
            suffix.set(i, nums.get(i+1) * suffix.get(i+1));
        }

        for (int i = 0; i < n; ++i) {
            res.set(i, prefix.get(i) * suffix.get(i));
        }

        return res;
    }

    public static void main(String[] args) {
        var nums = new ArrayList<>(List.of(1, 2, 3, 4));
        var prods = ProductOfArrayExceptSelf.doIt(nums);
        System.out.println("Expected [24,12,8,6], got " + prods);
    }
}
