package main

/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start

func findPoisonedDuration(timeSeries []int, duration int) int {

	// 设置区间左右边界
	leftMargin, rightMargin, total := timeSeries[0], timeSeries[0]+duration-1, 0

	for i := 1; i < len(timeSeries); i++ {
		// 落在区间内则延长右边界
		if timeSeries[i] <= rightMargin {
			rightMargin = timeSeries[i] + duration - 1
		} else {
			// 落在区间外则新建区间并统计旧区间的值
			total = total + rightMargin - leftMargin + 1
			leftMargin = timeSeries[i]
			rightMargin = timeSeries[i] + duration - 1
		}
	}

	return total + rightMargin - leftMargin + 1
}

// @lc code=end
