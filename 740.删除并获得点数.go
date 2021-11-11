package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {

	set := make(map[int]int)

	// nums = [2,2,3,3,3,4]
	// set = {2:2,3:3,4:1}

	for _, v := range nums {
		set[v]++
	}

	newNums := make([]int, 0)

	for v := range set {
		newNums = append(newNums, v)
	}

	// 将nums升序排列
	sort.Ints(newNums)

	dp := make([]int, len(newNums)+1)

	// newNums = [2,3,4]

	dp[0] = 0
	if len(newNums) == 0 {
		return dp[0]
	}

	dp[1] = newNums[0] * set[newNums[0]]
	if len(newNums) == 1 {
		return dp[1]
	}

	// 如果第二个元素 > 第一个元素+1 则取两者之和 否则取 两者的最大数
	if newNums[1] > newNums[0]+1 {
		dp[2] = dp[1] + newNums[1]*set[newNums[1]]
	} else {
		dp[2] = max_740(dp[1], newNums[1]*set[newNums[1]])
	}
	if len(newNums) == 2 {
		return dp[2]
	}

	// 动态规划状态转移
	// 假设第i位的数>前一位+1 则一起偷直接递加 否则需要比较 (前两位最大值)与(前一位值)的最大值

	for i := 3; i <= len(newNums); i++ {
		if newNums[i-1] > newNums[i-2]+1 {
			dp[i] = dp[i-1] + newNums[i-1]*set[newNums[i-1]]
		} else {
			dp[i] = max_740(dp[i-2]+newNums[i-1]*set[newNums[i-1]], dp[i-1])
		}
	}

	return max_740(dp[len(newNums)], dp[len(newNums)-1])
}

func max_740(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// func main() {
// 	fmt.Println(deleteAndEarn([]int{1, 6, 3, 3, 8, 4, 8, 10, 1, 3}))
// }
