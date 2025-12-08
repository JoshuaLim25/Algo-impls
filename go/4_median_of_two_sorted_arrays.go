package main

import "fmt"

// want a O(lgn) soln
func findMedianSortedArraysBRUTE(num1, num2 []int) float64 {
	median := 0.0
	l1, l2 := 0, 0
	merged := make([]int, 0, len(num1)+len(num2))
	for {
		if num1[l1] < num2[l2] {
			merged = append(merged, num1[l1])
			l1++
		} else {
			merged = append(merged, num2[l2])
			l2++
		}
		if l1 == len(num1) {
			merged = append(merged, num2[l2:]...)
			break
		}
		if l2 == len(num2) {
			merged = append(merged, num1[l1:]...)
			break
		}
	}
	fmt.Printf("%v\n", merged)
	l := len(merged)

	if l%2 == 0 { // even
		prev, next := merged[l/2-1], merged[l/2]
		fmt.Printf("%d, %v, %v\n", l, prev, next)
		// 1234
		median = float64(prev+next) / 2
	} else {
		median = float64(merged[l/2])
	}
	return median
}

// func findMedianSortedArrays(num1, num2 []int) int {
//
// }

func main() {
	n1 := []int{1, 3}
	n2 := []int{2}
	// nums1 := []int{1, 2}
	// nums2 := []int{3, 4}
	med := findMedianSortedArraysBRUTE(n1, n2)
	// med := findMedianSortedArraysBRUTE(nums1, nums2)
	fmt.Printf("%v\n", med)

}
