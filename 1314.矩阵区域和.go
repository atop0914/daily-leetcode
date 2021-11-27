package main

/*
 * @lc app=leetcode.cn id=1314 lang=golang
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	answer := make([][]int, m)

	for i := 0; i < m; i++ {
		ans := make([]int, n)
		for j := 0; j < n; j++ {
			ans[j] = func() int {
				sum := 0
				x, y, z, q := i-k, j-k, i+k, j+k
				if x < 0 {
					x = 0
				}
				if y < 0 {
					y = 0
				}
				if z > m-1 {
					z = m - 1
				}
				if q > n-1 {
					q = n - 1
				}

				for ii := x; ii <= z; ii++ {
					for jj := y; jj <= q; jj++ {
						sum += mat[ii][jj]
					}
				}
				return sum
			}()
			answer[i] = ans
		}
	}

	return answer
}

// @lc code=end

// func main() {
// 	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
// }
