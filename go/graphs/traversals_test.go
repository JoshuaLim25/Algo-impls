package graphs

import (
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput captures stdout during the call to f.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	return string(out)
}

func TestGraphPrintSearch(t *testing.T) {
	// Undirected graph:
	//    0
	//   / \
	//  1   2
	//   \ /
	//    3
	//    |
	//    4
	edges := [][2]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	nodes := BuildGraph(5, edges, false)

	type tc struct {
		name        string
		fn          func(*GNode, int) *GNode
		target      int
		wantContain string
	}

	tests := []tc{
		{"DFS_found", GraphDFS, 4, "Value found: 4"},
		{"BFS_found", GraphBFS, 4, "We found 4"},
		{"DFS_missing", GraphDFS, 99, ""},
		{"BFS_missing", GraphBFS, 99, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := strings.TrimSpace(captureOutput(func() {
				tt.fn(nodes[0], tt.target)
			}))
			if tt.wantContain == "" {
				if out != "" {
					t.Fatalf("expected no output, got %q", out)
				}
			} else {
				if !strings.Contains(out, tt.wantContain) {
					t.Fatalf("expected output to contain %q, got %q", tt.wantContain, out)
				}
			}
		})
	}
}
