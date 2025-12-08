/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice, finalProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		curProfit := prices[i] - minPrice
		if curProfit > finalProfit {
			finalProfit = curProfit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return finalProfit
}

func main() {

}
