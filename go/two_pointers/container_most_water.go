/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

package main

import (
	"fmt"
	"os"
)

func maxArea(height []int) int {
	// area calc: want to max(heights[l], height[r]) because this tells you the biggest possible height, then
	// NOTE: multiply this height with the width (*just* r-l and not r-l+1, since it's a span/dist and not a count)
	l, r := 0, len(height)-1
	curMax := 0
	for l < r { // don't want same line, 0 area
		fmt.Printf("l, r: (%d, %d)\n", l, r)
		fmt.Printf("Current max: %d", curMax)
		y := min(height[l], height[r]) // NOTE: MIN, b/c MAX would OVERFLOW
		x := r-l
		prod := x * y
		if prod > curMax {
			curMax = prod
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		// bad: doesnt handle the [x x l x x r] case, where l,r are final
		// how to keep a "good val" of r there? Not doing the below
		// advanceLeft := height[l+1]
		// advanceRight := height[r-1]
		// choice := max(advanceLeft, advanceRight)
		// if choice == advanceLeft {
		// 	l++
		// } else {
		// 	r--
		// }
	}
	fmt.Printf("l, r: (%d, %d)\n", l, r)

	return curMax
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}  // 49
	// height := []int{1, 1}  // 1
	area := maxArea(height)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: container_most_water.go:37: area=%+v\n", area)

}
