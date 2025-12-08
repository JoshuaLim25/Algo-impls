package graphs

import "fmt"

type GNode struct {
	Val int
	Adj []*GNode
}


func GraphDFS(start *GNode, needle int) *GNode {
	if start == nil {
		return nil
	}
	s := []*GNode{start}
	visited := map[*GNode]bool{start: true}

	for len(s) > 0 {
		nextUp := s[len(s)-1]
		s = s[:len(s)-1]

		if nextUp.Val == needle {
			fmt.Printf("Value found: %d\n", nextUp.Val)
		}

		for _, neighbor := range nextUp.Adj {
			// if neighbor != nil {  // not a tree
			if !visited[neighbor] {
				visited[neighbor] = true
				s = append(s, neighbor)
			}
		}
	}

	return nil
}

func GraphBFS(start *GNode, needle int) *GNode {
	if start == nil {
		return nil
	}
	q := []*GNode{start}
	head := 0
	visited := map[*GNode]bool{start: true}

	for head < len(q) {
		nextUp := q[head]
		head++

		if nextUp.Val == needle {
			fmt.Printf("We found %d\n", nextUp.Val)
		}

		for _, nb := range nextUp.Adj {
			if !visited[nb] {
				visited[nb] = true
				q = append(q, nb)
			}
		}

	}

	return nil
}

// BuildGraph returns a slice of n nodes (0..=n-1) wired according to edges.
// If directed is false, edges are added bidirectionally.
func BuildGraph(n int, edges [][2]int, directed bool) []*GNode {
	nodes := make([]*GNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &GNode{Val: i}
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		nodes[u].Adj = append(nodes[u].Adj, nodes[v])
		if !directed {
			nodes[v].Adj = append(nodes[v].Adj, nodes[u])
		}
	}
	return nodes
}
