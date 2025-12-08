/*
234. Palindrome Linked List - Easy

Given the head of a singly linked list, return true if it is a palindrome (reads the same backwards and forwards) or false otherwise.

Example 1:

Input: head = [1,2,2,1]
Output: true
*/

package main

// Works, but O(n) TC && NOT O(n) space.
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	repr := []int{}
	for head != nil {
		repr = append(repr, head.Val)
		head = head.Next
	}
	l, r := 0, len(repr)-1
	for l < r {
		if repr[l] != repr[r] {
			return false
		}
		l++
		r--
	}
	return true
}


func isPalindrome(head *ListNode) bool {
	once, twice := head, head
	for twice != nil && twice.Next != nil {
		once = once.Next
		twice = twice.Next.Next
	}
	// once here, once == middle of LL
	// Start reversing from ONCE
	// 1->2->3->3->2->1  |->  1->2->3<-1<-2<-3
	var prev *ListNode
	for once != nil {
		next := once.Next
		once.Next = prev
		prev = once
		once = next
	}
	// NOTE: prev will be the "last node"
	l, r := head, prev
	// for l != r {  // BAD:
	for r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true
}
