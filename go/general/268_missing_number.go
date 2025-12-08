/*
268. Missing Number - Easy

Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

All distinct, in range [0, n]
*/

func missingNumber(nums []int) int {
	// since [0,n], *should* be n+1 nums
	l := len(nums)
	expected := 0
	// for i := 1; i <= l; i++ {
	// 	expected += i // 0 + 1 + 2 + ...
	// }
	// better:
	expected := (l * (l+1))/2
	tmp := 0
	for i := range nums {
		tmp += nums[i]
	}
	return expected-tmp
}
