/*
141. Linked List Cycle
Easy

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/
package main

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	once, twice := head, head
	for once != nil && twice != nil && fast.Next != nil {
		// NOTE: You don't need twice.Next.Next != nil because:
		// If twice.Next.Next is nil, the assignment twice = twice.Next.Next just sets twice to nil
		// Next iteration, twice != nil fails and loop exits cleanly
		// also, fast.Next handles the once nil check: it's always scouting ahead
		once = once.Next
		twice = twice.Next.Next
		if once == twice {
			return true
		}
	}
	return false
}

