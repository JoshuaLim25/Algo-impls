/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:

Input: prices = [7,1,5,3,6,4]
*/

package main


func maxProfit(prices []int) int {
	// want to track minPrice, curMaxProfit
	if len(prices) <= 1 {
		return 0
	}
	minBuyPrice, totalProfit := prices[0], 0
	// 7,1,5,3,6,4
	for i := 1; i < len(prices); i++ {
		// profit = sell - buy. sell happens after buy
		curPrice := prices[i]
		profit := curPrice - minBuyPrice
		if profit > totalProfit {
			totalProfit = profit
		}
		if curPrice < minBuyPrice {
			minBuyPrice = curPrice
		}
	}
	return totalProfit
}

func main() {
	
}
