package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	/*
		爬到n层有两种方法：
		从n-1层爬上去
		从n-2层爬上去
		问题转为求爬到n-1层的方法+爬到n-2层的方法
		F(n) = F(n-1) + F(n-2)
	*/

	if n <= 2 {
		return n
	}

	// return climbStairs(n-1) + climbStairs(n-2)

	pre, end := 1, 2

	for i := 3; i <= n; i++ {
		tmp := pre + end
		pre = end
		end = tmp
	}

	return end
}

// @lc code=end
