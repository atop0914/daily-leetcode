package main

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// @lc code=end
