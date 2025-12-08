# Two Sum
- Stmt: Given an array of integers `nums` and an integer `target`, return the indices `i` and `j` such that `nums[i]` + `nums[j]` == `target` and `i` != `j`.
- You may assume that every input has *exactly one* pair of indices i and j that satisfy the condition. Return the answer with the smaller index first.

Example 1:
Input: 
nums = [3,4,5,6], target = 7
Output: [0,1]
Explanation: nums[0] + nums[1] == 7, so we return [0, 1].

Example 2:
Input: nums = [4,5,6], target = 10
Output: [0,2]


- Constraints:
2 <= nums.length <= 1000
-10,000,000 <= nums[i] <= 10,000,000
-10,000,000 <= target <= 10,000,000


# SOLN
```go

// Brute force.
// O(n^2)
func twoSum(nums []int, target int) (int, int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return 0, 0
}

// Hashtable, 2 passes.
// O(n)
func twoSum(nums []int, target int) (int, int) {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i := range nums {
		complement := target - nums[i]
		if j, ok := m[complement]; ok && j != i {
			return i, j
		}
	}
	return 0, 0
}
```
