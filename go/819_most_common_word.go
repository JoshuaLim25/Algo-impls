package main

import (
	"fmt"
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	var (
		freq          = make(map[string]int)
		bannedWords   = make(map[string]bool)
		mostFreqCount int
		finalWord     string
	)
	// 1. trim the input: punctuation and casing
	if len(paragraph) == 0 {
		return ""
	}
	re := regexp.MustCompile("[^a-zA-Z]+")
	trimmedParagraph := re.ReplaceAllString(paragraph, " ")
	fmt.Printf("%v\n", trimmedParagraph)
	// words := strings.Split(trimmedParagraph, " ")
	words := strings.Fields(trimmedParagraph)
	// filteredWords := words[:0] // reuse underlying array words
	filteredWords := []string{}

	// 1.b: get banned words out
	for _, w := range banned {
		bannedWords[w] = true
	}
	for i := range words {
		words[i] = strings.ToLower(words[i])
		if !bannedWords[words[i]] {
			filteredWords = append(filteredWords, words[i])
		}
	}

	fmt.Printf("words = %#v\n", filteredWords)
	fmt.Printf("words = %#v\n", words)

	// 2. count occurrences of each word in filtered words
	for _, w := range filteredWords {
		freq[w]++
	}

	fmt.Printf("%+v\n", freq)
	for k, occurrences := range freq {
		if occurrences > mostFreqCount {
			mostFreqCount = occurrences
			finalWord = k
		}
	}

	return finalWord
}

func main() {
	// s := "Bob hit a ball, the hit BALL flew far after it was hit."
	// banned := []string{"hit"}
	s := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}
	res := mostCommonWord(s, banned)
	fmt.Printf("%+v\n", res)

}
