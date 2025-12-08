/*
53. Maximum Subarray - Medium

Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/

package main
/*
// fukcing kadanes
def maxSubArray(self, nums: List[int]) -> int:            
	res = nums[0]
	total = 0

	for n in nums:
		if total < 0:
			total = 0

		total += n
		res = max(res, total)

	return res
*/

func sumArray(s []int) []int {
	sum := 0
	for i := range s {
		sum += s[i]
		s[i] = sum
	}
	return s
}
func maxSubArrayBOTCHED(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sums := make([]int, 0, len(nums))
	sums = sumArray(nums)
	l := 0
	maximum := nums[0]
	for r := 1; r < len(nums); r++ {
		curMax := sums[r] - sums[l]
		if curMax > maximum {
			maximum = curMax
		} 
	}

}

func kadanes(nums []int) int {
	res := nums[0]
	total := 0

	for _, num := range nums {
		if total < 0 {
			total = 0
		}
		total += num
		res = max(res, total)
	}
	return res
}
