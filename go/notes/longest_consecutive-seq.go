package main 

import(
	"fmt"
	"os"
)

type Set map[int]struct{}

func LongestConsecutiveSeq(nums []int) int {
	var longest int
	set := Set {}
	// go thru lsit, check if each elem present. If it is, skip, if not insert
	for _, num := range nums {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
		}
	}
	
	fmt.Fprintf(os.Stderr, "DEBUG[6]: longest-consequtive-seq.go:12: set=%+v\n", set)
	longest = len(set)
	return longest
}

func main() {
	nums1 := []int {0,3,7,2,5,8,4,6,0,1}
	nums2 := []int {100,4,200,1,3,2}
	longest1 := LongestConsecutiveSeq(nums1)
	fmt.Fprintf(os.Stderr, "DEBUG[4]: longest-consequtive-seq.go:26: longest1 should be 9, is %+v\n", longest1)
	longest2 := LongestConsecutiveSeq(nums2)
	fmt.Fprintf(os.Stderr, "DEBUG[5]: longest-consequtive-seq.go:28: longest2 should be 4, is %+v\n", longest2)
}
