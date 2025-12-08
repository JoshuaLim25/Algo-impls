/*
206. Reverse Linked List
Easy
Topics
premium lock iconCompanies

Given the head of a singly linked list, reverse the list, and return the reversed list.

*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode { // return new head
	// 1 -> 2 -> 3
	// 3 -> 2 -> 1
	var prev *ListNode  // nil
	for head != nil {
		// ORDER MATTERS: MODIFY NEXT, THEN OG HEAD
		next := head.Next  // 1. save access to res of list
		head.Next = prev   // 2. switch the link (reverse)
		prev = head        // 3. move prev so next iteration works
		head = next        // 4. same argument
	}
	return prev
	// essentially, prev is always behind by one, and once head becomes nil
	// prev will store the last elt (new head)
}
