/*
78. Subsets - Medium

Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package com.algos.backtracking;

import java.util.ArrayList;
import java.util.List;

public class Subsets78 {
    List<List<Integer>> res = new ArrayList<>();
    List<Integer> curSet = new ArrayList<>();
    List<Integer> nums;

    public List<List<Integer>> subsets(List<Integer> nums) {
        // backtracking involves (1) goals, (2) constraints, (3) choices made and taken back
        // nums = [1,2,3];
        if (nums == null || nums.size() == 0) return res;
        this.nums = nums;

        backtrack(0);
        return res;
    }

    private void backtrack(int idx) {
        // res.add(curSet);   // BAD: need the temp copy since it's mutated by reference (i.e., via closure; other backtrack calls)
        res.add(new ArrayList<>(curSet));

        for (int i = idx; i < nums.size(); ++i) {
            curSet.add(nums.get(i));
            backtrack(i+1);
            curSet.removeLast();
        }
    }

    public static void main(String[] args) {
        var nums = List.of(1,2,3);
        var s = new Subsets78();
        System.out.println(s.subsets(nums));
    }
}
