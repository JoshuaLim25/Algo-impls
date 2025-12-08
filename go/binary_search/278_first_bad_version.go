/*
278. First Bad Version
Easy
Topics
premium lock iconCompanies

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.
*/

package main

func firstBadVersion(n int) int {
	// NOTE: n nonzero
	l, r := 1, n  // TODO: don't forget return version+1
	potentialFirst := 0
	for l <= r {
		mid := l + (r-l)/2
		// isBadVersion(num) returns true if bad, false if not
		// [f f f f t t t t t]
		// all the ones to right will be false
		if isBadVersion(mid) {  // t, first will be to the left
			potentialFirst = mid
			r = mid - 1
		} else { // f, will be to the right
			l = mid + 1
		}
	}
	return potentialFirst
}
