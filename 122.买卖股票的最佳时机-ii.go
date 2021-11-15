package main

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
/*
func maxProfit(prices []int) int {

	// 思想：我可以针对每个prices[i]算出截止第 i day,我所可以获取的最大的收益
	// 其中我需要保存截止第 i day 最低的成本 minCost,及阶段性的最大收益maxProfit
	// 起始位置需要判断一下prices[0]与price[1],因为我肯定需要挑一个相对低点买入。假设我可以知道前两天的股票价格 这题不考虑亏损的情况


	if len(prices) == 1 {
		return 0
	}

	if len(prices) == 2 && prices[1] < prices[0] {
		return 0
	}

	maxProfit, hasStock, minCost, days := 0, false, prices[0], len(prices)

	// dp := make([]int, len(prices))

	// 遍历价格数组 如果我的最小花费为0 表示我可以买入当天的股票；如果我的最小花费不为0 表示我可以卖出当天的股票
	for i := 1; i < days; i++ {

		// 假设手中没有股票
		if !hasStock {
			minCost = min_122(minCost, prices[i])
			hasStock = true
		} else {
			// 有股票
			if prices[i] > minCost {
				// 卖出股票 收益递增
				if i == days-1 {
					maxProfit += prices[i] - minCost
					hasStock = false
					continue
				}
				maxProfit += max_122(prices[i], prices[i+1]) - minCost
				hasStock = false

				minCost = prices[i+1]
			}

		}

	}

	return max_122(maxProfit, maxProfit_121(prices))
}

func max_122(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min_122(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProfit_121(prices []int) int {
	maxProfit, minCost, days := 0, prices[0], len(prices)

	for i := 1; i < days; i++ {
		minCost = min_122(minCost, prices[i])
		maxProfit = max_122(prices[i]-minCost, maxProfit)
	}

	return maxProfit
}
*/

func maxProfit(prices []int) int {
	maxProfit := 0

	// 看了评论才会的 当天卖出 当天买入 利润为0 核心思想就是只要第二天的价格大于第一天的就卖出 获利递增
	for i := 1; i < len(prices); i++ {

		profit := prices[i] - prices[i-1]

		if profit > 0 {
			maxProfit += profit
		}

	}

	return maxProfit
}

// @lc code=end
