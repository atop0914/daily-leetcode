package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	/*
		最大子序列和
		假设dp数组为截止到nums index和的最大值
		dp[0] = nums[0]
		dp[1] = max(nums[1], nums[1]+dp[0])
		dp[n] = max(nums[n], nums[n]+dp[n-1])

	*/
	if len(nums) == 1 {
		return nums[0]
	}

	length := len(nums)
	// dp := make([]int, length+1)
	subMax := 0
	max := math.MinInt32
	left, right, subLeft, subRight := 0, 0, 0, 0 // 最大子数组下标

	for i := 0; i < length; i++ {
		if subMax > 0 { // 前最大子序列和大于0 所以右下标往后移
			subRight++
			subMax += nums[i]
		} else {
			subMax = nums[i]
			subLeft, subRight = i, i
		}
		if subMax >= max {
			max = subMax
			left, right = subLeft, subRight
		}
	}
	fmt.Println(nums[left : right+1])
	return max
}

func max_53(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
