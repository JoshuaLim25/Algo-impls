package main

import (
	"fmt"
	"strings"
)
func frequencySort(s string) string {
	// Note: case must be preserved, A and a are distinct
	// init freq map
	freq := make(map[rune]int)
	repr := []rune(s)
	var sb strings.Builder
	for _, r := range repr {
		freq[r]++
	}
	var (
		biggest  int
		nextRune rune
	)

	fmt.Printf("%v\n", freq)
	for len(freq) > 0 {
		fmt.Printf("%v\n", biggest)
		for k, v := range freq {
			if v > biggest {
				fmt.Printf("%v\n", v)
				biggest = v
				nextRune = k
			}
		}
		for biggest > 0 {
			sb.WriteRune(nextRune)
			biggest--
		}
		// biggest == 0 for next iter
		delete(freq, nextRune)
	}
	res := sb.String()
	return res
	// keep ranging over map
}

func main() {
	s := "tree" // "eert"
	res := frequencySort(s)
	fmt.Printf("%v\n", res)
	
}
