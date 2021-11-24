package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	min, length := math.MaxInt32, len(matrix)
	dp := matrix[length-1]

	// 从倒数第二层开始遍历往上
	for i := length - 2; i >= 0; i-- {
		tmp := make([]int, length)
		for j := 0; j < length; j++ {
			if j >= 1 && j <= length-2 {
				tmp[j] = min_931(matrix[i][j]+dp[j], min_931(matrix[i][j]+dp[j+1], matrix[i][j]+dp[j-1]))
			} else if j == 0 {
				tmp[j] = min_931(matrix[i][j]+dp[j+1], matrix[i][j]+dp[j])
			} else {
				tmp[j] = min_931(matrix[i][j]+dp[j-1], matrix[i][j]+dp[j])
			}
		}
		dp = tmp
	}

	for _, v := range dp {
		if min > v {
			min = v
		}
	}

	return min
}

func min_931(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
