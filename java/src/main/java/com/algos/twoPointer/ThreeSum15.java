/*
15. 3Sum - Medium

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set *must not contain duplicate triplets*.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
*/

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;


public class ThreeSum15 {
    public static List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        int targetSum = 0;
        // Input: nums = [-1,0,1,2,-1,-4]; want 0
        // sorted nums = [-4,-1,-1,0,1,2]
        //                 i  j        k  = -3 != 0
        //                 i     j     k  = -3 != 0
        //                 i       j   k  = -2 != 0
        //                 i         j k  = -1 != 0, and so on. advance i and continue
        // idea: first pointer i ranges from 0..nums.length - 2; i.e.., stops 3 elts behind end
        // -2 since two pointers in the way :)

        // NOTE: O(nlgn) but needs to happen:
        Arrays.sort(nums);

        for (int i = 0; i < nums.length - 2; ++i) {
            // NOTE: THIS IS THE SAUCE TO REMEMBER (i.e., check for dupes from where you came)
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            // init j, k as l, r ptrs respectively.
            int j = i + 1, k = nums.length - 1;
            while (j < k) {  // NO overlap, since one cond says "i != j, i != k, and j != k" (i.e., distinct)
                int sum = nums[i] + nums[j] + nums[k];
                if (sum == targetSum && i != j && i != k && j != k) {
                    res.add(new ArrayList<>(List.of(nums[i], nums[j], nums[k])));
                    // here, you need to do the same check for updating and seeing if any other solns in this window [i..end]
                    // it also needs to be a while, since you keep j++ing until you get non-dupe (ie [...i..j.....k] -> [...i....j...k] -> etc.
                    while (j < k && nums[j] == nums[j + 1]) { // for l, look to its right
                        j++;
                    }
                    while (j < k && nums[k] == nums[k - 1]) { // for r, look to its left
                        k--;
                    }
                    // then this is what you'd do if you didn't care about dupes.
                    // i.e., you found a solution, so advance each ptr and see if you can find more in the zone
                    j++;
                    k--;
                } else if (sum < targetSum) {
                    j++;
                } else if (sum > targetSum) {
                    k--;
                }
            }
        }
        return res;
    }

    public static void main(String[] args) {
        int[] nums1 = new int[]{-1, 0, 1, 2, -1, -4};
        int[] nums2 = new int[]{-100, -70, -60, 110, 120, 130, 160};
        System.out.println(String.format("[[-1,-1,2],[-1,0,1]] == %s", ThreeSum15.threeSum(nums1)));
        System.out.println(String.format("[[-100,-60,160],[-70,-60,130]] == %s", ThreeSum15.threeSum(nums2)));


    }
}


