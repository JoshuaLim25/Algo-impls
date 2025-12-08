package main

import "fmt"

type GNode struct {
	Val any
	Adj []*GNode
}

type Node struct {
	Val any
	left *Node
	right *Node
}

func DFSRec(start *Node, needle int) bool {
	if start == nil {
		return false
	}
	if start.Val == needle {
		return true
	}
	if start.left != nil {
		return DFSRec(start.left, needle)
	}

	if start.right != nil {
		return DFSRec(start.right, needle)
	}
	return false
}

func DFSIter(start *Node, needle int) *Node {
	if start == nil {
		return nil
	}
	s :=[]*Node{start}

	for len(s) > 0 {
		//pop
		nextUp := s[len(s)-1]
		s = s[:len(s)-1]

		// process
		if nextUp.Val == needle {
			fmt.Printf("Data found: %d", nextUp.Val)
			return nextUp
		}

		// push children
		if nextUp.right != nil {
			s = append(s, nextUp.right)
		}
		if nextUp.left != nil {
			s = append(s, nextUp.left)
		}
	}
	return nil
}

func GraphDFSIter(start *GNode, needle int) *GNode {
	if start == nil {
		return nil
	}

	s := []*GNode{start}
	visited := map[*GNode]bool{ start: true }

	for len(s) > 0 {
		// pop
		nextUp := s[len(s)-1]
		s = s[:len(s)-1]

		// process 
		if nextUp.Val == needle {
			fmt.Printf("Found data %d\n", nextUp.Val)
			return nextUp
		}

		// add to q
		for _, nb := range nextUp.Adj {
			if !visited[nb] {
				visited[nb] = true
				s = append(s, nb)
			}
		}
	}
	return nil
}

func BFS(start *Node, needle int) *Node {
	if start == nil {
		return nil
	}
	q := []*Node{start}
	head := 0

	for head < len(q) {
		// pop
		nextUp := q[head]
		head++
		
		// process
		if nextUp.Val == needle {
			fmt.Printf("Data found: %d", nextUp.Val)
			return nextUp
		}

		// add children
		if nextUp.left != nil {
			q = append(q, nextUp.left)
		}
		if nextUp.right != nil {
			q = append(q, nextUp.right)
		}
	}
	return nil

}

func main() {

}

