package main

import "math"

/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
// func maxScoreSightseeingPair(values []int) int {
// 	maxScore, lenght := 0, len(values)

// 	for i := 0; i < lenght; i++ {
// 		for j := i + 1; j < lenght; j++ {
// 			maxScore = max_1014(maxScore, values[i]+values[j]+i-j)
// 		}
// 	}

// 	return maxScore
// }
func maxScoreSightseeingPair(values []int) int {

	/*
		values[i] + values[j] - (j-i) 可以得到i-j的观光得分
		期望 values[i] + values[j] 越大越好
		j-i 越小越好

		dp[i] = values[i] + values[j] + i - j
		dp[i] = values[i]+i + values[j]-j


		values[j]-j是固定值, 求values[i]+i的最大值
		遍历数组的时候维护values[i]+i的值记为 max
	*/

	max, ans := values[0], math.MinInt32

	for j := 1; j < len(values); j++ {
		ans = max_1014(max+values[j]-j, ans)
		max = max_1014(max, values[j]+j)
	}
	return ans
}

func max_1014(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
