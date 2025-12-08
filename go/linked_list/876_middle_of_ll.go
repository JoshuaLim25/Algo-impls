/*
876. Middle of the Linked List - Easy

Given the head of a singly linked list, return the middle node of the linked list.
**If there are two middle nodes, return the second middle node.**
*/

package main


type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	once := head  // *ListNode
	twice := head
	for once != nil && twice != nil && twice.Next != nil {
		once = once.Next
		twice = twice.Next.Next
	}
	return once
}

