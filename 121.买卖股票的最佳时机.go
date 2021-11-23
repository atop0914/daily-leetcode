package main

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit_121(prices []int) int {
	maxProfit, minCost, days := 0, prices[0], len(prices)

	for i := 1; i < days; i++ {
		minCost = min_121(minCost, prices[i])
		maxProfit = max_121(prices[i]-minCost, maxProfit)
	}

	return maxProfit
}

func max_121(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min_121(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
