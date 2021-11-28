package main

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid[0]), len(grid)
	dp := make([]int, m)
	dp[0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < n; i++ {
		tmp := grid[i]
		for j := 0; j < m; j++ {
			if j == 0 {
				dp[j] += tmp[j]
			} else {
				dp[j] = min_64(dp[j-1]+tmp[j], dp[j]+tmp[j])
			}
		}
	}
	return dp[m-1]
}

func min_64(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

// func main() {
// 	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
// }
