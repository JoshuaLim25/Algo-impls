package main

import (
	"fmt"
	"os"
)

// copy of avgs that looked for []float64, also find the max

// Stmt: Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
// Property: contiguous subarrays, len == k. `n` == len(input_array), `k` == len(subarray)
// Edge cases: nil array, dont care neg 0 pos elems (add indiscrimately, div by n). n < k?
// Input: Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
// Output: Output: [2.2, 2.8, 2.4, 3.6, 2.8]

func averageSubarray(data []int, k int) []float64 {
	if len(data) == 0 {
		return nil
	}

	var curSum int
	// this assumes len(data) >= k
	for i := range k {
		curSum += data[i]
	}
	sums := []int{curSum}

	l, r := 0, k  // data[:k] accumulated
	for ; r < len(data); r++ {
		curSum += data[r]
		curSum -= data[l]
		sums = append(sums, curSum)
		l++
	}
	fmt.Fprintf(os.Stderr, "DEBUG[1]: avg-contiguous-subarray_2.go:24: sums=%+v\n", sums)

	avgs := make([]float64, 0, len(sums))
	for _, sum := range sums {
		avg := float64(sum)/float64(k)
		avgs = append(avgs, avg)
	}
	fmt.Fprintf(os.Stderr, "DEBUG[3]: avg-contiguous-subarray_2.go:38: avgs=%+v\n", avgs)

	maxAvg := 0.0
	for _, avg := range avgs {
		maxAvg = max(maxAvg, avg)
	}
	fmt.Fprintf(os.Stderr, "DEBUG[1]: max-avg-contiguous-subarray_2.go:45: maxAvg=%+v\n", maxAvg)


	return avgs
}

func main() {
	data := []int{1,12,-5,-6,50,3}
	fmt.Fprintf(os.Stderr, "DEBUG[1]: avg-contiguous-subarray_2.go:46: data=%+v\n", data)
	k := 4
	averages := averageSubarray(data, k)
	fmt.Fprintf(os.Stderr, "DEBUG[4]: avg-contiguous-subarray_2.go:44: averages=%+v\n", averages)

	var totalAvg float64
	for _, avg := range averages {
		totalAvg += avg
	}
	fmt.Fprintf(os.Stderr, "DEBUG[1]: avg-contiguous-subarray_2.go:52: totalAvg=%+v\n", totalAvg)

}
