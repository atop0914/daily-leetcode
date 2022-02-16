package main

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	res := this.queue1[0]

	this.queue2 = this.queue1[1:]

	this.queue1, this.queue2 = this.queue2, this.queue1

	return res
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.queue1[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.queue1) == 0 {
		return true

	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
