/*
322. Coin Change - Medium

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
Return the *fewest* number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an *infinite number of each kind of coin*.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/

/*
Goal: use an []int to **count combinations** that sum to `amount`
Edge cases:
1. Coints unique
2. Answer fits in a signed 32-bit int
3. Strictly positive
*/

package main 

func coinChangeII(amount int, coins []int) int {
	if len(coins) == 0 && amount > 0 {
		return 0
	}
	// Answers Q of: starting with coin[i], how many ways can we make curAmt?
	var numWaysToMakeAmount func(i, curAmt int) int
	numWaysToMakeAmount = func(i, curAmt int) int {
		// base case for skip action
		if i == len(coins) {  // meaning there are no coins to use
			if curAmt == 0 {
				return 1  // 1 way to make 0 with no coins
			} else {
				return 0
			}
		}
		take := 0
		// Decision to take the ith coin, resulting in curAmt-coins[i]
		if coins[i] <= curAmt {
			// NOTE: acts as its own base case, since curAmt is decreased,
			// and eventually the outer conditonal will fail
			take = numWaysToMakeAmount(i, curAmt - coins[i])
		}
		// Decision to skip the ith coin and move on to the next
		// TODO: how to bound this? Need some base case.
		// Notice that it always calls i++. When does this need to end?
		skip := numWaysToMakeAmount(i+1, curAmt)
		return skip + take
	}
	return numWaysToMakeAmount(0, amount)
}
