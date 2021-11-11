package main

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {

	/*

		每次爬1层或者2层 爬到 len(cost)有 climbStairs(n) = climbStairs(n-1)+climbStairs(n-2)

		爬到n层有两种方法：
		①
		先踏上第i-2级台阶（最小总花费dp[i-2]），再直接迈两步踏上第i级台阶（花费cost[i]），最小总花费dp[i-2] + cost[i]；

		②
		先踏上第i-1级台阶（最小总花费dp[i-1]），再迈一步踏上第i级台阶（花费cost[i]），最小总花费dp[i-1] + cost[i]；

		状态转移方程：
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]


		初始条件：

		最后一步踏上第0级台阶，最小花费dp[0] = cost[0]。

		最后一步踏上第1级台阶有两个选择：
		①
		可以分别踏上第0级与第1级台阶，花费cost[0] + cost[1]
		②
		也可以从地面开始迈两步直接踏上第1级台阶，花费cost[1]
		最小值dp[1] = min(cost[0] + cost[1], cost[1]) = cost[1]


	*/

	length := len(cost)

	dp := make([]int, length)

	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < length; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}

	return min(dp[length-1], dp[length-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
