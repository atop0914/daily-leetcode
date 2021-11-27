package main

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {

	/*
		以m=3, n=2为例：抵达(3,2)有两种方法
		① 从(3,1)往右
		② 从(2,2)往下
		A(3,2)=A(3,1)+A(2,2)
		A(3,1)=1
		A(2,2)=2
		A(m,n)=A(m-1,n)+A(m,n-1)
	*/

	dp := make([][]int, m+1)
	seq := make([]int, n+1)

	for i := range seq {
		if i != 0 {
			seq[i] = 1
		}
	}
	dp[1] = seq

	for i := 2; i < m+1; i++ {
		seq := make([]int, n+1)
		for j := 1; j < n+1; j++ {
			if j == 1 {
				seq[1] = 1
			}
			seq[j] = seq[j-1] + dp[i-1][j]
		}
		dp[i] = seq
	}

	return dp[m][n]
}

// @lc code=end
