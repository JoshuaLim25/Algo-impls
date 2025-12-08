package lol

import (
	"fmt"
)

// input := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}; len(input) = 9
func AvgSubarray(a []int, k int) (avgs []float64) {
	// given [ nums ] -> [ avgs ]
	low := 0
	sum := 0
	for end := range a {
		sum += a[end]
		if end >= k-1 {
			avg := float64(sum) / float64(k)
			avgs = append(avgs, avg)
			sum -= a[low]
			low++
		}
	}
	if low == len(a)-k-1 {
		fmt.Printf("Correct: %v\n", low)
	} else {
		fmt.Printf("NT: %v\n", low)
	}

	return avgs
}

func BinarySearch[T ~int | ~float64 | ~string](data []T, needle T) (idx int, found bool) {
	if len(data) == 0 || data == nil {
		return idx, found
	}
	low, high := 0, len(data)-1
	for low <= high {
		mid := (high + low) / 2
		if data[mid] > needle {
			high = mid - 1
		} else if data[mid] < needle {
			low = mid + 1
		} else {
			return mid, !found
		}
	}
	return idx, found
}

func MakeSlice[T ~int | ~float64 | ~string](dx int, dy int) *[][]T {
	res := make([][]T, dx)
	for i := range res {
		res[i] = make([]T, dy)
	}
	return &res
}
