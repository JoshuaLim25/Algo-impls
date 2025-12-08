/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		} else {
			return l1
		}
	}
	dummy := ListNode{}
	cur := &dummy
	// back and forth merging
	for l1 != nil && l2 != nil {
		lesser := min(l1.Val, l2.Val)
		if lesser == l1.Val {
			cur.Next = l1
			l1 = l1.Next
		} else if lesser == l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else if l1 == l2 {
			// pick at random
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// check if one still has other stuff
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return dummy.Next
}

func mergeTwoListsRec(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1 = mergeTwoListsRec(l1.Next, l2)
		return l1
	}
	l2 = mergeTwoListsRec(l1, l2.Next)
	return l2
}
