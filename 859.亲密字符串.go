package main

/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {

	// 长度不一样肯定不是亲密字符串
	if len(s) != len(goal) {
		return false
	}

	// 字符串完全一样
	if s == goal {
		// 判断字符串中有无重复的字符
		set := make(map[byte]int, 0)
		for pos := 0; pos < len(goal); pos++ {
			set[goal[pos]]++
		}

		for _, v := range set {
			if v >= 2 {
				return true
			}
		}
	} else {
		// 字符串不一样
		des1 := make([]byte, 0, 2)
		des2 := make([]byte, 0, 2)

		for pos := 0; pos < len(goal); pos++ {

			// 集合长度超过2退出
			if len(des1) > 2 {
				return false
			}

			// 字节不相等就把字节分别压入两个集合
			if s[pos] != goal[pos] {
				des1 = append(des1, s[pos])
				des2 = append(des2, goal[pos])
			}
		}

		// 集合长度不为2退出
		if len(des1) != 2 {
			return false
		}

		if des1[0] == des2[1] && des1[1] == des2[0] {
			return true
		}
	}
	return false
}

// @lc code=end
// func main() {
// 	fmt.Println(buddyStrings("abac", "aba"))
// }
