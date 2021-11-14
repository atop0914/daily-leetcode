package main

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	/*
		假设第求以第i为结尾最大乘积为dp[i]
		假设nums[i]为负数；则期望dp[i-1]也为负数且越小越好
		假设nums[i]为正数；则期望dp[i-1]为正数且越大越好
		状态转移方程:
		dp[i] = max(dp[i], dp[i-1]*nums[i], dm[i-1]*nums[i])
		三者求最大值
	*/
	maxProduct, minProduct, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		maxP, minP := maxProduct, minProduct

		maxProduct = max_152(max_152(nums[i]*maxP, nums[i]), minP*nums[i])

		minProduct = min_152(min_152(nums[i]*minP, nums[i]), nums[i]*maxP)

		ans = max_152(maxProduct, ans)
	}

	return ans
}

func max_152(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min_152(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end
