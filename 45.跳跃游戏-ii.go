package main

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start

func jump(nums []int) int {

	/*
		贪婪算法，我们每次在可跳范围内选择可以使得跳的更远的位置。
		[2,3,1,4,1]

		① nums[0]可以触达 nums[1]、nums[2]由于 nums[1]最远触达nums[4] nums[2]最远触达nums[3] 所以下一次跳到nums[1] step++

		② nums[1]可以触达 nums[2]、nums[3]、nums[4] nums[4]到达右边界直接跳到nums[4] step++

		step=2
	*/

	var steps, end int // 总步数、当前所在位置的索引

	maxPosition, length := 0, len(nums) // 下一次可以跳到的最大右边界

	for i := 0; i < length-1; i++ {
		maxPosition = max_45(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func max_45(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// func main() {
// 	fmt.Println(jump([]int{2, 3, 1, 4, 1}))
// }
