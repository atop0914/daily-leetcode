package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l3 := new(ListNode)
	po3 := l3

	var carry int // 向前进位

	for l1 != nil || l2 != nil || carry != 0 {

		po3.Next = new(ListNode)
		po3 = po3.Next

		var num1, num2 int

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		po3.Val = (num1 + num2 + carry) % 10
		carry = (num1 + num2 + carry) / 10

	}

	return l3.Next

}

// @lc code=end
