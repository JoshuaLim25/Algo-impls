package main

import (
	"fmt"
	"os"
)

func BinarySearch(haystack []int, needle int) int {
	l, r := 0, len(haystack)-1
	for l <= r {
		mid := l + (r-l)/2
		if haystack[mid] == needle {
			return mid
		} else if haystack[mid] > needle {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}



func main() {
	data := []int { 2, 3, 5, 6, 7, 9, 100}
	target := 3
	idx := BinarySearch(data, target)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: binary_search.go:21: idx=%+v\n", idx)

}
