package main

/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start

/* 直接递归会超时
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
}
*/

func tribonacci(n int) int {

	first, second, third := 0, 1, 1

	if n < 3 {
		switch n {
		case 0:
			return first
		case 1:
			return second
		case 2:
			return third
		}
	}

	// generate nums in array
	for i := 3; i <= n; i++ {
		tmp := first + second + third
		first = second
		second = third
		third = tmp
	}

	return third
}

// @lc code=end

// func main() {
// 	fmt.Println((tribonacci(25)))
// }
