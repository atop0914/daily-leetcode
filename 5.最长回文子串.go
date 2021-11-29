package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	//dp(i,k)=dp(i+1,k-1)&&s(i)==s(k)
	var ans string
	dp := make([][]bool, len(s))

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for k := 0; k <= i; k++ {
			dp[i][k] = s[i] == s[k] && (i-1 < k+1 || dp[i-1][k+1])
			if dp[i][k] && i-k+1 > len(ans) {
				ans = s[k : i+1]
			}
		}
	}

	return ans
}

// @lc code=end
