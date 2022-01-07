package main

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Stack struct {
	s []int
}

func newStack() *Stack {
	return &Stack{
		s: make([]int, 0),
	}
}

func (s *Stack) Push(node int) {
	s.s = append(s.s, node)
}

func (s *Stack) Pop() int {
	if len(s.s) == 0 {
		return -1
	}

	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return res
}

func isPalindrome(head *ListNode) bool {

	p, q := head, head

	stack := newStack()

	for p != nil {
		stack.Push(p.Val)
		p = p.Next
	}

	for q != nil {
		v := stack.Pop()
		if v != q.Val {
			return false
		}
		q = q.Next
	}

	return true
}

// @lc code=end

// func main() {
// 	a := new(ListNode)
// 	b := new(ListNode)
// 	c := new(ListNode)
// 	d := new(ListNode)

// 	a.Val = 1
// 	b.Val = 2
// 	c.Val = 2
// 	d.Val = 1

// 	a.Next = b
// 	b.Next = c
// 	c.Next = d
// 	d.Next = nil

// 	fmt.Println(isPalindrome(a))

// }
