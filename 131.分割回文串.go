package main

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start

var res [][]string

// var path []string

func partition(s string) [][]string {

	res = make([][]string, 0)
	path := make([]string, 0)

	dfs_131(s, 0, path, &res)

	return res
}

func dfs_131(str string, index int, path []string, res *[][]string) {
	// 1.确定截止条件
	if index == len(str) {
		// lst := make([]string, len(path))
		// copy(lst, path)
		*res = append(*res, path)
		return
	}

	// 2.候选节点
	for i := 1; i <= len(str)-index; i++ {
		x := str[:index+i]
		// 2.1 候选节点是否满足筛选条件
		// 看候选节点是否满足回文要求即reverse结果非reverse结果一致
		if x == reverse(x) {
			path = append(path, x)
			dfs_131(str[index+i:], index, path, res)
			path = path[:len(path)-1]
		}
	}

}

// 对一个字符串做翻转
func reverse(x string) string {
	runes := []rune(x)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end
