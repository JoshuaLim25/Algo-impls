package main

import (
	"fmt"
)

type Node struct {
	data int
	left *Node
	right *Node
}

func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%v\n", n.data)  // Visit
	preOrder(n.left)
	preOrder(n.right)
}

func inOrder(n *Node) {
	if n == nil {
		return
	}
	inOrder(n.left)
	fmt.Printf("%v\n", n.data)  // Visit
	inOrder(n.right)
}

func postOrder(n *Node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Printf("%v\n", n.data)  // Visit
}

func insert(n *Node, data int) *Node {
	if n == nil {
		return &Node { data: data }
	}

	if data < n.data {
		n.left = insert(n.left, data)
	} else if data > n.data {
		n.right = insert(n.right, data)
	}

	return n
}



func DFSRec(n *Node, needle int) *Node {
	if n == nil {
		return nil
	}
	if n.data == needle {
		fmt.Printf("Needle found: %d\n", n.data)
		return n
	}
	if left := DFSRec(n.left, needle); left != nil {
		return left
	}
	return DFSRec(n.right, needle)
}

func DFSIter(n *Node, needle int) bool {
	if n == nil {
		return false
	}
	s := []*Node{n}
	for len(s) > 0 {
		// pop
		top := s[len(s)-1]
		s = s[:len(s)-1]

		// process
		if top.data == needle {
			fmt.Printf("Data found: %d\n", top.data)
			return true
		}

		// repop w/children
		if top.right != nil {
			s = append(s, top.right)
		}
		if top.left != nil {
			s = append(s, top.right)
		}

	}

	return false
}

func BFS(n *Node, needle int) {
	if n == nil {
		return
	}
	q := []*Node{n}
	// O(1) amortized dequeue, **no reslice**
	head := 0              // new
	// for len(q) != 0 {   // og
	for head < len(q) {    // new
		// Pop
		// nextUp := q[0]  // og
		// q = q[1:]       // og
		nextUp := q[head]  // new
		head++             // new
		// Process
		if nextUp.data == needle {
			fmt.Printf("Data found: %d\n", nextUp.data)
			return
		}
		if nextUp.left != nil {
			q = append(q, nextUp.left)
		}
		if nextUp.right != nil {
			q = append(q, nextUp.right)
		}
	}
}


func main() {
	root := &Node{
		data: 1,
		left: nil,
		right: nil,
	}

	insert(root, 2)
	insert(root, 3)
	insert(root, 4)
	insert(root, 5)
	insert(root, 6)
	insert(root, 7)
	insert(root, 8)
	insert(root, 9)
	insert(root, 10)
	insert(root, 11)
	insert(root, 12)
	insert(root, 13)
	insert(root, 14)
	insert(root, 15)

	preOrder(root)
	inOrder(root)
	postOrder(root)

	// BFS(root, 6)
	// good := DFSRec(root, 6)
	DFSIter(root, 6)
}
