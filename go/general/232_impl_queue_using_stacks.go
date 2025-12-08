/*
Interface: push, peek, pop, and empty
*/
package main

type MyQueue struct {
	Stack []int
	Queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		Stack: []int{},
		Queue: []int{},
	}
}

func (q *MyQueue) Pop() int {
	if len((*q).Stack) == 0 {
		stackToQueue((*q).Stack, (*q).Queue)
	}
	top := (*q).Stack[len((*q).Stack)-1]
	(*q).Stack = (*q).Stack[:len((*q).Stack)-1]
	return top
}

func (q *MyQueue) Push(elt int) {
	if len((*q).Stack) == 0 {
		stackToQueue((*q).Stack, (*q).Queue)
	}
	(*q).Stack = append((*q).Stack, elt)
}

func (q *MyQueue) Peek() int {
	if len((*q).Stack) == 0 {
		stackToQueue((*q).Stack, (*q).Queue)
	}
	return (*q).Queue[len((*q).Queue)-1]
}

func (q *MyQueue) Empty() bool {
	return len((*q).Queue) + len((*q).Stack) == 0
}


func stackToQueue(s, q []int) {
	for len(s) > 0 {
		top := s[len(s)-1]
		s = (s)[:len(s)-1]
		q = append(q, top)
	}
}
