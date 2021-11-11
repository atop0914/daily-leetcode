package main

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
/*
var slen int
var exist *bool

func isSubsequence(s string, t string) bool {
	path := make([]byte, 0)
	exist := new(bool)
	slen = len(s)
	dfs_392(s, t, path, exist)
	return *exist
}

func dfs_392(s string, t string, path []byte, exist *bool) {
	// 确定边界条件
	if len(s) == 0 || len(t) == 0 {
		if len(path) == slen {
			*exist = true
			return
		}
	}

	if len(s) != 0 && len(t) != 0 {
		for i := 0; i < len(t); i++ {
			// 候选节点
			if t[i] == s[0] {
				path = append(path, s[0])
				dfs_392(s[1:], t[i+1:], path, exist)
			}
		}
	}
}
*/

// 双指针法
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	// 双指针分别指向s,t的第0位
	pos1, pos2 := 0, 0

	for pos1 < len(s) && pos2 < len(t) {
		// 匹配到同时后移
		if s[pos1] == t[pos2] {
			pos1++
			pos2++
		} else {
			// 未匹配到则pos2后移
			pos2++
		}
	}

	return pos1 == len(s)
}

// @lc code=end

// func main() {
// 	b := isSubsequence("axc", "ahbgdc")
// 	fmt.Println(b)
// }
