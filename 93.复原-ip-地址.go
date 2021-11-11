/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
package main

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	ress := make([]string, 0)
	dfs_93(s, -1, 1, res, &ress)
	return ress
}

// dfs深度优先遍历
func dfs_93(str string, index, level int, res []string, ress *[]string) {
	// 截止条件
	if level == 5 || index == len(str)-1 {
		if level == 5 && index == len(str)-1 {
			*ress = append(*ress, (res[0] + "." + res[1] + "." + res[2] + "." + res[3]))
		}
		return
	}

	// 候选人
	for i := 1; i < 4; i++ {

		x := substr(str, index+1, i)

		// 2.1 筛选
		if index <= len(str)-1 && isNormal(x) {
			res = append(res, x)
			dfs_93(str, index+i, level+1, res, ress)
			res = res[:len(res)-1]
		}
	}
}

func isNormal(str string) bool {
	checkInt, _ := strconv.Atoi(str)
	if len(str) > 1 && str[0] == '0' {
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// func main() {
// 	s := restoreIpAddresses("0000")
// 	fmt.Println(s)
// }

// @lc code=end
