package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit_309(prices []int) int {

	// dp[i][0] 表示第i天没有未持有股票的最大收益
	// dp[i][1] 表示第i天持有股票的最大收益

	length := len(prices)

	if length <= 1 {
		return 0
	}

	dp := make([][]int, length+1)

	dp[0] = make([]int, 2)
	dp[0][0] = 0
	dp[0][1] = math.MaxInt32

	dp[1] = make([]int, 2)
	dp[1][0] = 0
	dp[1][1] = -prices[0]

	for i := 2; i <= length; i++ {
		dp[i] = make([]int, 2)
		// 今天未持有的最大利润是以下两者中的最大值：
		// 1. 昨天未持有，今天未持有
		// 2. 昨天持有，今天卖出
		dp[i][0] = max_309(dp[i-1][1]+prices[i-1], dp[i-1][0])

		// 如果今天持有，其中就有一种可能就是今天买的，那么就必须考虑前天的状态
		// 因为前天的状态会影响今天的行为：
		// 1. 昨天持有，今天持有
		// 2. 昨天未持有，此时只能考虑前天未持有状态。否则违反冷冻期规定
		// 而昨天未持有，前天未持有的收益为 dp[i-2][0]
		dp[i][1] = max_309(dp[i-1][1], dp[i-2][0]-prices[i-1])
	}

	return dp[length][0]
}

func max_309(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
