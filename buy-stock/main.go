package main

import "fmt"

/*
给定一个数组代表股票每天的价格，请问只能买卖一次的情况下，最大化利润是多少？日期不重叠的情况下，可以买卖多次呢？
输入: {100, 80, 120, 130, 70, 60, 100, 125}
1）只能买一次：65（60 买进，125 卖出）
2）可以买卖多次: 115（80买进，130卖出；60 买进，125卖出）
输出买卖的序列和最大利润
*/

type Pair struct {
	i, j int
}

func main() {
	prices := []int{100, 80, 120, 130, 70, 60, 100, 125}

	i, j := buy_and_sell_once(prices)

	fmt.Println(prices[i], prices[j])

	res := buy_and_sell_multiple_times(prices)
	for _, index := range res {
		fmt.Print(prices[index], " ")
	}
	fmt.Println()
}

// buy_and_sell_once return (i, j) that maximize (p[j] - p[i]).
func buy_and_sell_once(prices []int) (int, int) {
	// min_index[j]: the index i that i <= j and p[i] is the smallest.
	min_index := make([]int, len(prices))

	// max_profit[k]: the index pair (i,j) that i < j <= k and p[j] - p[i] is the biggest.
	max_profit := make([]Pair, len(prices))

	// max_profit[k+1] = max(p[k+1] - min_index[k+1], max_profit[k])

	min_index[0] = 0
	max_profit[0] = Pair{}

	for i := 1; i < len(prices); i++ {
		// calculate min_index
		if prices[i] < prices[min_index[i-1]] {
			min_index[i] = i
		} else {
			min_index[i] = min_index[i-1]
		}

		// calculate max_profit
		previous_pair := max_profit[i-1]
		if prices[i]-prices[min_index[i-1]] > prices[previous_pair.j]-prices[previous_pair.i] {
			max_profit[i] = Pair{min_index[i-1], i}
		} else {
			max_profit[i] = max_profit[i-1]
		}
	}

	last := max_profit[len(prices)-1]
	return last.i, last.j
}

func buy_and_sell_multiple_times(prices []int) []int {
	res := make([]int, 0)
	bought := false
	buy_candidate_index := 0
	sell_candidate_index := 0

	for k := 1; k < len(prices); k++ {
		switch bought {
		case false:
			// find a low price to buy
			if prices[k] <= prices[buy_candidate_index] {
				// wait for another lower price
				buy_candidate_index = k
			} else { // prices[k] > prices[buy_candidate_index]
				// buy at buy_candidate_index
				bought = true
				sell_candidate_index = k
			}
		case true:
			// finder a high price to sell
			if prices[k] >= prices[sell_candidate_index] {
				if k < len(prices)-1 {
					// wait for the next higher price.
					sell_candidate_index = k
				} else {
					// end: have to sell
					res = append(res, buy_candidate_index, k)

					return res
				}
			} else { // prices[k] < prices[sell_candidate_index]
				// deal: sell at sell_candidate_index
				res = append(res, buy_candidate_index, sell_candidate_index)
				buy_candidate_index = k
				bought = false
			}
		}
	}

	return res
}
