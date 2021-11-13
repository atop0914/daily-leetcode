package main

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
	dp := make([]int, length+1)
	max := nums[0]

	dp[0] = nums[0]
	dp[1] = max_53(nums[1], nums[1]+dp[0])

	for i := 1; i < length; i++ {
		dp[i] = max_53(nums[i], nums[i]+dp[i-1])
		if dp[i] >= max {
			max = dp[i]
		}
	}

	return max
}

func max_53(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
