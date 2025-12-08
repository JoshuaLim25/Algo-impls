/*
46. Permutations - Medium

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 */

// package com.algos.backtracking;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class Permutations46 {
    List<List<Integer>> res = new ArrayList<>();
    List<Integer> perms = new ArrayList<>();
    Set<Integer> seen = new HashSet<>();
    int[] nums;

    public List<List<Integer>> permute(int[] nums) {
        if (nums.length == 0)
            return this.res;
        this.nums = nums;

        backtrack();
        return this.res;
    }

    private void backtrack() {
        if (perms.size() == nums.length) {
            res.add(new ArrayList<>(perms));
            return;
        }

        for (int i = 0; i < nums.length; ++i) {
            if (seen.contains(i))
                continue;
            seen.add(i);
            perms.add(nums[i]);
            backtrack();
            seen.remove(i);
            perms.removeLast();
        }
    }

    public static void main(String[] args) {
        var empty = new int[] {};
        System.out.println(Arrays.stream(empty)
                .boxed()
                .collect(Collectors.toList()));

        Permutations46 soln = new Permutations46();
        soln.nums = new int[] { 1, 2, 3 };

        soln.permute(soln.nums);
        System.out.println(soln.res);
    }
}
