/*
The general approach to dp problems is determining, for a given state, most of the time irrespetive of how you got there:
1. What matters? If it's a max sum, if it's a position, etc
2. Does the previous step you took matter (e.g., only moving horiz/vertically in a mxn grid)?

Generally, consider DP when:
1. Greedy approach won't work (e.g., find shortest path by taking shortest cost at ech step)
2. You want to do a bunch of similar computations and also cache the results of a prior computation (this is dp bread and butter)
3. The problem itself deals with:
- Combinations/permutations/number of ways
-

STMT:
You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?


https://www.youtube.com/watch?v=YBSt1jYwVfU
here, i is your position (what stop you are on). dp[i] is your soln. since the question asks for number of ways, this is `int dp[i]`. If it were a predicate (T/F is it possible) it would be `bool dp[i]`
*/

package dynamic_programming

func FibDP(n int) int {
	dp := make([]int, n+2)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ { // until it equals n
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs(n int) int {
	// n=2 |-> dp(1) + dp(0), n=3 |-> dp(2) + dp(1)
	dp := make([]int, n+2)
	dp[0], dp[1] = 1, 1             // for no steps or one step, only one way
	for pos := 2; pos <= n; pos++ { // until it equals n
		dp[pos] = dp[pos-1] + dp[pos-2]
	}
	return dp[n]
}
