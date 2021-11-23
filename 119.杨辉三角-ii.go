package main

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {

	numRows := rowIndex + 1

	dp := make([][]int, numRows)

	// 输入: numRows = 5
	// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

	dp[0] = []int{1}

	if numRows <= 1 {
		return dp[rowIndex]
	}

	dp[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1

		for j := 1; j < i; j++ {
			tmp[j] = dp[i-1][j-1] + dp[i-1][j]
		}

		tmp[i] = 1
		dp[i] = tmp
	}

	return dp[rowIndex]
}

// @lc code=end
