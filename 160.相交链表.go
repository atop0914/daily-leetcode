package main

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	// m1 := make(map[*ListNode]bool)

	// for headA != nil {
	// 	m1[headA] = true
	// 	headA = headA.Next
	// }

	// for headB != nil {
	// 	if m1[headB] == true {
	// 		return headB
	// 	}
	// 	headB = headB.Next
	// }

	// return nil

	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}

	}
	return p1
}

// @lc code=end
