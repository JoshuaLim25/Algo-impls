package main

import(
	"fmt"
	"os"
)

type Letters map[rune]struct{}

func ValidAnagram(s, t string) bool {
	letters := Letters {}
	for _, char := range s {
		letters[char] = struct{}{}
	}
	for _, char := range t {
		if _, present := letters[char]; !present {
			return false
		}
	}
	return true
}


func main() {
	s, t := "racecar", "racecar"
	var(
		a int  // 0
		b = 2
		f = 3.14
	)
	fmt.Fprintf(os.Stderr, "DEBUG[9]: valid_anagram.go:29: f=%+v\n", f)
	fmt.Fprintf(os.Stderr, "DEBUG[7]: valid_anagram.go:27: a=%+v\n", a)
	fmt.Fprintf(os.Stderr, "DEBUG[8]: valid_anagram.go:29: b=%+v\n", b)
	// s, t := "rat", "car"  // false
	// s, t := "anagram", "nagaram"  // true
	ValidAnagram(s, t)
	fmt.Fprintf(os.Stderr, "DEBUG[6]: valid_anagram.go:26: t=%+v\n", t)
	fmt.Fprintf(os.Stderr, "DEBUG[5]: valid_anagram.go:26: s=%+v\n", s)
	fmt.Fprintf(os.Stderr, "DEBUG[4]: valid_anagram.go:28: ValidAnagram(s, t)=%+v\n", ValidAnagram(s, t))

}
