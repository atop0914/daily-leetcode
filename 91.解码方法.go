package main

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {

	length := len(s)
	step := make([]int, length+1)

	step[0] = 1

	for i := 1; i <= length; i++ {
		if s[i-1] != '0' {
			step[i] += step[i-1]
		}

		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			step[i] += step[i-2]
		}
	}

	return step[length]
}

// @lc code=end
