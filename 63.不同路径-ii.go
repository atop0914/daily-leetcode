package main

import "fmt"

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	/*
		思考 A(m,n)=A(m-1,n)+A(m,n-1) 当u(m,n)=0时
		A(m,n)=0 当u(m,n)=1时(有障碍的时候)
	*/

	m, n := len(obstacleGrid[0]), len(obstacleGrid)
	dp := make([]int, m)

	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m-1]
}

// @lc code=end

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
