/*
207. Course Schedule - Medium

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

- for pair [x,y], x is the course you want to take in [0,numCourses), y is prereq
- are pairs unique? yes

Input: n: 5, so [0,4); prerequisites: [[0,1]]
*/

package main

func canFinish(numCourses int, prereq [][]int) bool {
	// Init step
	reqMap := make(map[int][]int)
	for i := range numCourses {
		reqMap[i] = []int{}
	}
	for _, pair := range prereq {
		course, prereq := pair[0], pair[1]
		reqMap[course] = append(reqMap[course], prereq)
	}
	visited := make([]bool, numCourses) // stores courses along the current DFS path
	return true
}
