/*
100. Same Tree - Easy

Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/

package main

func isSameTreeBROKEN(p, q *TreeNode) bool {
	// WHY: Else branches will never allow recursive calls to happen
	// p == q only works when both are nil, but fails when one is nil and other isn't

	// root
	if p == q {
		return true
	} else {
		return false
	}
	// values
	if p.Val == q.Val {
		return true
	} else {
		return false
	}
	// children
	if p.Left == q.Left {
		return true
	} else {
		return false
	}
	if p.Right == q.Right {
		return true
	} else {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTreeALSOBROKEN(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// if p != nil || q != nil {  // BAD:
	if p == nil || q == nil {  // works since the nil, nil check above
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, p.Left) && isSameTree(p.Right, p.Right) 
}


func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// if p != nil || q != nil {  // BAD: we already checked if they were empty. We would expect both to be non-nil.
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


func isSameTree(p, q *TreeNode) bool {
    // Both nil
    if p == nil && q == nil {
        return true
    }
    // One nil, one not
    if p == nil || q == nil {
        return false
    }
    // Different values
    if p.Val != q.Val {
        return false
    }
    // Recursively check both subtrees
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
