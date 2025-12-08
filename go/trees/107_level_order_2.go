package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
		// return [][]int{}  // why
	}
	res := [][]int{}
	// res := make([][]int, 0) // why
	q := []*TreeNode{root}
	for len(q) > 0 {
		l := len(q)
		children := make([]int, 0, l)
		for i := range q {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
			// works b/c if left there, right not only one append,
			// and also no weird stuff like [50, 0]
			children = append(children, q[i].Val)
		}
		q = q[l:]
		res = append(res, children)
	}
	return res
}

func levelOrderBottom(root *TreeNode) [][]int {
	tmp := levelOrder(root)
	reversed := [][]int{}
	for i := len(tmp)-1; i >= 0; i-- {
		reversed = append(reversed, tmp[i])
	}
	return reversed
}


func main() {

}
