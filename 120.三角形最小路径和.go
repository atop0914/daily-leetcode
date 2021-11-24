package main

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {

	ans, length := 0, len(triangle)

	dp := triangle[length-1] // 令dp等于最后一行数组

	for i := length - 2; i >= 0; i-- {
		tmp := make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			tmp[j] = min_120(triangle[i][j]+dp[j], triangle[i][j]+dp[j+1])
		}
		dp = tmp
	}

	ans = dp[0]

	return ans
}

func min_120(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
