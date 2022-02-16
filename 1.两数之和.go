package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range nums {
		if i, ok := m[target-value]; ok {
			return []int{index, i}
		} else {
			m[value] = index
		}
	}

	return []int{}
}

// @lc code=end
