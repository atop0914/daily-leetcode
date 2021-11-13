package main

/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	max, maxNow := -2<<31, -2<<31
	min, minNow := 2<<31, 2<<31
	maxVal := -2 << 31
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
		maxVal = Max(maxVal, nums[i])

		maxNow = Max(nums[i], maxNow+nums[i])
		max = Max(max, maxNow)

		minNow = Min(nums[i], minNow+nums[i])
		min = Min(min, minNow)
	}

	if maxVal <= 0 { // 如果数组中的数都小于0 返回其中那个最大的
		return maxVal
	}

	return Max(max, total-min)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
