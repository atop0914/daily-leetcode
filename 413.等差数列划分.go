package main

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	ans, add := 0, 0

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			add++
			ans += add
		} else {
			add = 0
		}

	}

	return ans
}

// @lc code=end
