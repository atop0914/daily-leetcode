package main

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start

/*

 */

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	// 遍历元素
	max := 0

	for index, v := range nums {
		if index > max {
			continue
		}

		if index+v >= max {
			max = index + v
		}

	}

	return max >= len(nums)-1
}

// @lc code=end
