package com.algos.search;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Result {
    int idx;
    boolean present;

    public Result(int idx, boolean present) {
        this.idx = idx;
        this.present = present;
    }

    @Override
    public String toString() {
        return "com.algos.search.Result @idx " + this.idx + ", present = " + this.present;
    }
}

class RotatedSortedArray {
    public Result search(List<Integer> nums, int target) {
        int l = 0, r = nums.size() - 1;
        while (l <= r) {  // <=
            int mid = l + (r-l) / 2;
            if (target == nums.get(mid)) {
                return new Result(mid, true);
            }
            // on left [4567] or right [123]?
            // check if on left
            if (nums.get(l) <= nums.get(mid)) {
                // check if in bounds
                if (nums.get(l) <= target && target < nums.get(mid)) {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            } else { // on right
                if (nums.get(mid) < target && target <= nums.get(r)) {
                    l = mid + 1;
                } else {
                    r = mid - 1;
                }
            }
        }
        return new Result(-1, false);
    }

    public static void main(String[] args) {
        List<Integer> nums = new ArrayList<>(Arrays.asList(4, 5, 6, 7, 1, 2, 3));
        int target = 3;
        RotatedSortedArray r = new RotatedSortedArray();
        Result res = r.search(nums, target);
        System.err.println("DEBUG[1]: com.algos.search.RotatedSortedArray.java:38: res=" + res);
    }
}
