package dynamic_programming

// func CoinChange(coins []int, amount int) (best int) {
// 	numWaysToMake := func(i int, curr_amt int) int {
// 		if i == 0 {
// 			if curr_amt == 0 {
// 				return 1
// 			} else {
// 				return 0
// 			}
// 		}
// 		take := 0
// 		if coins[i] <= curr_amt {
// 			take = numWaysToMake(i, curr_amt-coins[i])
// 		}
// 		skip := numWaysToMake(i+1, curr_amt)
// 		return take + skip
// 	}
// 	return best
// }

type KnapsackData struct {
	Weights  []int
	Profits  []int
	Capacity int
}

// given a bag with capacity m, need to find the items that maximize the profit gained form putting any number of items into the bag (no dup)
func Knapsack(bag *KnapsackData) [][]int {
	res := make([][]int, len(bag.Weights))
	for i := range len(bag.Weights) {
		res[i] = make([]int, bag.Capacity+1)
	}

	for i := range len(bag.Weights) {
		for w := range bag.Capacity {
			if i == 0 || w == 0 {
				res[i][w] = 0
			} else if bag.Weights[i] <= w {
				res[i][w] = max(res[i-1][w], res[i-1][w-bag.Weights[i]]+bag.Profits[i])
			} else {
				res[i][w] = res[i-1][w]
			}
		}
	}
	return res
}
