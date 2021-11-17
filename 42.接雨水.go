package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max_42(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max_42(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min_42(leftMax[i], rightMax[i]) - h
	}
	return
}

func min_42(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max_42(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
