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

	min := math.MaxInt32

	for i := 0; i < len(matrix); i++ {
		tmpMin := matrix[0][i] // 从第一行第i个元素开始累加
		flg := i
		for j := 1; j < len(matrix); j++ {
			ttm := 0
			if flg >= 1 && flg <= len(matrix)-2 {
				ttm = min_931(matrix[j][flg-1], min_931(matrix[j][flg], matrix[j][flg+1]))
				if ttm == matrix[j][flg-1] {
					flg = flg - 1
				} else if ttm == matrix[j][flg+1] {
					flg = flg + 1
				}
			} else if flg == 0 {
				ttm = min_931(matrix[j][flg+1], matrix[j][flg])
				if ttm == matrix[j][flg+1] {
					flg = flg + 1
				}
			} else if flg == len(matrix)-1 {
				ttm = min_931(matrix[j][flg-1], matrix[j][flg])
				if ttm == matrix[j][flg-1] {
					flg = flg - 1
				}
			}
			tmpMin += ttm
		}

		if tmpMin < min {
			min = tmpMin
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
