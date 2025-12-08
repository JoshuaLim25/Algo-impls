package com.algos.twoPointers;

import java.util.Arrays;

public class ContainerWithMostWater11 {
    public static int maxAmount(int[] heights) {
        // idea: have two pointers on either end.
        // [ 1 8 6 2 5 4 8 3 7 ]
        //   ^               ^
        // care about the max area, which is w x h
        int l = 0, r = heights.length - 1;
        int maxArea = 0;
        while (l < r) {  // just Bisualize. if you crossed/intercepted l, r then invalid area (width == 0)
            int width = r-l;  // span of distance or discrete count of elts?
            int height = Math.min(heights[l], heights[r]);  // min, since the shorter one is the weak link
            int area = width * height;

            if (height == heights[l]) l++;
            else if (height == heights[r]) r--;
            maxArea = Math.max(maxArea, area);
        }
        return maxArea;
    }

    public static void main(String[] args) {
        var nums = new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7 };
        System.out.println(String.format("ContainerWithMostWater11.maxAmount(%s) should be 49: %d", nums.toString(), ContainerWithMostWater11.maxAmount(nums)));
    }
}
