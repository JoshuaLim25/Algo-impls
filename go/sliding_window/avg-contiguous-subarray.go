package sliding_window

// Stmt: Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
// Property: contiguous subarrays, len == k. `n` == len(input_array), `k` == len(subarray)
// Edge cases: nil array, dont care neg 0 pos elems (add indiscrimately, div by n). n < k?
// Input: Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
// Output: Output: [2.2, 2.8, 2.4, 3.6, 2.8]

func AvgSubarrayBad(data []int, k int) (output []float64) {
	if len(data) < k {
		return nil
	}
	sum := 0
	for high := k; high < len(data); high++ {
		sum = 0
		for low := high - k; low < high; low++ {
			sum += data[low]
		}
		avg := float64(sum) / float64(k)
		output = append(output, avg)
	}
	return output
}

func AvgSubarrayBrute(data []int, k int) (output []float64) {
	if len(data) < k {
		return nil
	}
	sum := 0
	for high := k - 1; high < len(data); high++ {
		sum = 0
		for low := high - (k - 1); low <= high; low++ {
			sum += data[low]
		}
		avg := float64(sum) / float64(k)
		output = append(output, avg)
	}
	return output
}

func AvgSubarraySW(data []int, k int) (output []float64) {
	start := 0 // this is advanced once window is of size k
	sum := 0   // this accumulates start..end of window
	for end := range len(data) {
		sum += data[end]
		if end > len(data)-k-1 { // end of the valid sliding winodw, -1 for index
			avg := float64(sum) / float64(k)
			output = append(output, avg)
			sum -= data[start]
			start++
		}
	}
	return output
}

type ListNode struct {
	Val  any
	Next *ListNode
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) {
//
// }
