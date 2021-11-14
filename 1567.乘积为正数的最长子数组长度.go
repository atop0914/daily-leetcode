package main

/*
 * @lc app=leetcode.cn id=1567 lang=golang
 *
 * [1567] 乘积为正数的最长子数组长度
 */

// @lc code=start
func getMaxLenV1(nums []int) int {

	length := len(nums) // 数组长度

	positive, negative := make([]int, length), make([]int, length)

	if nums[0] > 0 {
		positive[0] = 1
	} else if nums[0] < 0 {
		negative[0] = 1
	}

	maxLength := positive[0]

	for i := 1; i < length; i++ {
		if nums[i] == 0 {
			positive[i] = 0
			negative[i] = 0
		} else if nums[i] > 0 {
			positive[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				negative[i] = negative[i-1] + 1
			} else {
				negative[i] = 0
			}
		} else if nums[i] < 0 {
			negative[i] = positive[i-1] + 1

			if negative[i-1] > 0 {
				positive[i] = negative[i-1] + 1
			} else {
				positive[i] = 0
			}
		}
		maxLength = max_1567(maxLength, positive[i])
	}

	return maxLength
}

func getMaxLen(nums []int) int {
	/*
		任意正积区段必然不含0，因此以零为分界点。

		对于任意一个非零区段，其中的正积区段来源于以下几种情况：
		1）若其中负数有偶数个，则是区段本身；
		2）若其中负数有奇数个，则是最左边负数右侧、或是最右边负数左侧（两者取最长）；
		https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/solution/shuang-zhi-zhen-by-wanglongjiang-9u9y/
	*/
	ans, n := 0, len(nums)

	left, right := 0, 0 // 双指针起始位置

	neg := 0 // 负数的数量

	for left < n {
		for right < n && nums[right] != 0 {
			if nums[right] < 0 {
				neg++
			}
			right++
			if neg%2 == 0 {
				ans = max_1567(ans, right-left)
			}
		}

		for left+1 < right {
			if nums[left] < 0 {
				neg--
			}
			left++
			if neg%2 == 0 {
				ans = max_1567(ans, right-left)
			}
		}

		for right < n && nums[right] == 0 {
			right++
		}
		left = right
		neg = 0
	}

	return ans
}

func max_1567(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
