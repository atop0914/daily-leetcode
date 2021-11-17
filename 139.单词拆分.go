package main

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {

	length := len(s)
	wordExist, dp := make(map[string]bool), make([]bool, length+1)

	for _, v := range wordDict {
		wordExist[v] = true
	}

	dp[0] = true

	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordExist[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[length]
}

// @lc code=end
