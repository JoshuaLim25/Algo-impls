package main

import(
	"fmt"
	"os"
)

// horiz or vert. No diagonal, directly adjacent
type Board [][]byte

func BoardInit(rows, cols int) Board {
	b := make([][]byte, rows)
	for row := range b {
		b[row] = make([]byte, cols)
	}
	return b
}

type Pos struct {
	x int
	y int
}

// 1. at each step, check our current position/cell. If the value there is the same as the letter we are on in the word, advance.
// what it means to advance: first mark that we've visited it, then check everywhere horizontally and vertically && check visited.
// continue recusively.
func Exists(board Board, word string) bool {
	exist := true
	valid := true
	curLetter := 0
	curPos := Pos{x: 0, y: 0}
	visited := []Pos{}
	letters := []byte(word)
	for valid {
		if board[curPos.x][curPos.y] == letters[curLetter] {
			
		}
		valid = false
	}
	

	return !exist
}

func main() {
	b := BoardInit(3, 4)
	for _, row := range b {
		fmt.Printf("Contents of row: %v\n", row)
			}
	b1 := Board {
		{ 'A','B','C','E' },
		{ 'S','F','C','S' },
		{ 'A','D','E','E' },
	}
	fmt.Fprintf(os.Stderr, "DEBUG[3]: word_search.go:29: b1=%+v\n", b1)
	word1 := "ABCCED"
	fmt.Fprintf(os.Stderr, "DEBUG[1]: word_search.go:34: word1=%+v\n", word1)
	Exists(b1, word1)
}
