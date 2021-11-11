package main

import "fmt"

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {

	/*
		动态规划

		偷第i间房屋 此时金额 dp[i]=dp[i-2]+nums[i]

		偷第i-1间房屋 此时金额 dp[i-1]=dp[i-3]+nums[i-1]

		问题转为求
		max((dp[i-2]+nums[i]), (dp[i-3]+nums[i-1]))

	*/

	dp := make([]int, len(nums)+1)
	n := len(nums)

	dp[0] = 0
	if n == 0 {
		return dp[0]
	}

	dp[1] = nums[0]
	if n == 1 {
		return dp[1]
	}

	dp[2] = max(nums[0], nums[1])
	if n == 2 {
		return dp[2]
	}

	for i := 3; i <= n; i++ {
		dp[i] = max((dp[i-2] + nums[i-1]), (dp[i-3] + nums[i-2]))
	}

	fmt.Println(dp)

	return max(dp[n-1], dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
