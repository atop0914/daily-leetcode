package main

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {

	// 因此每天交易结束后只可能存在手里有一支股票或者没有股票的状态。
	// dp[i][0]表示第i天结尾手中没有股票的最大收益
	// dp[i][1]表示第i天结尾手中有一只股票的最大收益

	length := len(prices)

	dp := make([][]int, length+1)
	dp[0], dp[1] = make([]int, 2), make([]int, 2)
	dp[0][0] = 0
	dp[0][1] = 0

	// 第一天手中没有股票收益为0  第一天手中有股票说明第一天买入了股票此时收益为-prices[0]
	dp[1][0] = 0
	dp[1][1] = -prices[0]

	for i := 2; i < length+1; i++ {
		dp[i] = make([]int, 2)
		// 手中没有股票
		// ① dp[i-1][1]+当天卖出的价格 前一天有股票今天卖出
		// ② dp[i-1][0] 前一天没有股票 今天也没有
		dp[i][0] = max_714(dp[i-1][1]+prices[i-1]-fee, dp[i-1][0])

		// 手中有股票
		// ① dp[i-1][0]-当天买入的价格 前一天没有有股票 今天买入股票
		// ② dp[i-1][1] 前一天有股票 今天没有进行卖出
		dp[i][1] = max_714(dp[i-1][0]-prices[i-1], dp[i-1][1])
	}

	// 最后一天肯定不希望手中还有股票
	// 所以取最后一天手中没有股票的情况
	return dp[length][0]
}

func max_714(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
