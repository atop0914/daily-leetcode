package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	list3 := new(ListNode) // 新链表的头结点
	p3 := list3            // 指针指向头结点

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p3.Next = list1
			list1 = list1.Next
		} else {
			p3.Next = list2
			list2 = list2.Next
		}
		p3 = p3.Next
	}

	if list1 != nil {
		p3.Next = list1
	}

	if list2 != nil {
		p3.Next = list2
	}

	return list3.Next
}

// @lc code=end
