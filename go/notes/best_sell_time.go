package main

import(
	"fmt"
	"os"
)

type Days struct {
	buy int
	sell int
}

func BestSellTime(prices []int) int {
	totalProfit := 0
	// if profit generated (sellPrice-buyPrice) is <=0, no updates to curProfit
	/// Brute: go through list. In a nested loop, calculate each profit.
	m := make(map[int]Days)  // { profit: buy+sell days }
	for buy := range prices {
		for sell := buy + 1; sell < len(prices); sell++{
			buyPrice, sellPrice  := prices[buy], prices[sell]
			profit := sellPrice - buyPrice
			m[profit] = Days{
				buy: buyPrice,
				sell: sellPrice,
			}
		}
	}
	for k, v := range m {
		fmt.Fprintf(os.Stderr, "DEBUG[1]: best_sell_time.go:22: k, v = %+v, %+v\n", k, v)
		if k > totalProfit {
			totalProfit = k
		}
	}
	fmt.Fprintf(os.Stderr, "DEBUG[3]: best_sell_time.go:35: totalProfit=%+v\n", m[totalProfit])
	return totalProfit
}

func main() {
	// prices := []int{7,1,5,3,6,4}
	prices := []int{7,6,4,3,1}

	profit := BestSellTime(prices)
	fmt.Fprintf(os.Stderr, "DEBUG[2]: best_sell_time.go:38: profit=%+v\n", profit)

}
