package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	path := make([]int, 0)
	set := make(map[int]int) // 已遍历的数组将位置保存与set集合中

	for index, value := range nums {
		// 每次遍历都去map中查找有没有匹配元素 如果有的话就直接退出 没有的话把元素存放于map中
		if val, ok := set[target-value]; ok {
			path = append(path, []int{val, index}...)
			break
		} else {
			set[value] = index
		}
	}

	return path
}

// @lc code=end

func main() {
	fmt.Println(twoSum([]int{9, 2, 3, 11, 1, 2, 3, 4, 5}, 6))
}
