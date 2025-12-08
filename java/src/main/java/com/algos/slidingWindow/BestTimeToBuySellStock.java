/*
121. Best Time to Buy and Sell Stock - Easy

You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 */

package com.algos.slidingWindow;

import java.util.ArrayList;
import java.util.List;

public class BestTimeToBuySellStock {
    public static int maxProfit(List<Integer> prices) {
        // want to buy low, sell high
        if (prices == null || prices.size() <= 1) return 0;
        int buyPrice = prices.get(0);
        int curProfit = 0;

        for (int i = 1; i < prices.size(); ++i) {
            int sellPrice = prices.get(i);
            int profit = sellPrice - buyPrice;

            // simple as
            // WHY: you're reassigning to the same variable even if you wrap it in an if check. You want to know if the next day you *could* potentially buy (sellPrice) is lower than your current buyPrice (sellPrice, since it could also just not be less and you'd just keep the same), and if so reassign to buyPrice. Both cases involve manipulating/mutating (or not mutating) buyPrice.
            curProfit = Math.max(curProfit, profit);
            buyPrice = Math.min(buyPrice, sellPrice);
        }

        return curProfit;
    }

    public static void main(String[] args) {
        var prices1 = new ArrayList<>(List.of(7, 1, 5, 3, 6, 4));
        System.out.println(String.format("Should return 5: %d", BestTimeToBuySellStock.maxProfit(prices1)));
        var prices2 = new ArrayList<>(List.of(7,6,4,3,1));
        System.out.println(String.format("Should return 0: %d", BestTimeToBuySellStock.maxProfit(prices2)));
    }
}
