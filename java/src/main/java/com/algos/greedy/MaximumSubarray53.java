/*
53. Maximum Subarray
Solved
Medium
Topics
conpanies iconCompanies

Given an integer array nums, find the

with the largest sum, and return its sum.

 

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
*/


public class MaximumSubarray53 {
    public static int maximumsubarray(int[] nums) {
        if (nums == null || nums.length == 0) return 0;
        int total = 0;
        int sum = nums[0];

        for (int n : nums) {
            if (total < 0) {
                total = 0;
            }
            total += n;
            sum = Math.max(sum, total);
        }
        return sum;
    }
    public static void main(String[] args) {
        int[] nums = {-2,1,-3,4,-1,2,1,-5,4};
        System.out.println(String.format("6 == %d", MaximumSubarray53.maximumsubarray(nums)));


        int[] nums2 = {1};
        System.out.println(String.format("1 == %d", MaximumSubarray53.maximumsubarray(nums2)));

    }
}
