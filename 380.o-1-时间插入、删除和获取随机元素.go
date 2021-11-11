package main

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	contentSet   map[int]int
	contentArray []int
}

// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		contentSet:   make(map[int]int),
// 		contentArray: make([]int, 0),
// 	}
// }

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.contentSet[val]; ok {
		return false
	}
	this.contentArray = append(this.contentArray, val)
	this.contentSet[val] = len(this.contentArray) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.contentSet[val]; ok {
		// 先到map中根据val查找val所在的value(数组中的index)

		// 交换array中的值 将index索引对应的值与array末尾中的值交换
		this.contentArray[index], this.contentArray[len(this.contentArray)-1] = this.contentArray[len(this.contentArray)-1], this.contentArray[index]

		// 更新map中最后一个元素的value
		this.contentSet[this.contentArray[index]] = index

		delete(this.contentSet, val)
		this.contentArray = this.contentArray[:len(this.contentArray)-1]

		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {

	num := rand.Intn(len(this.contentSet))

	return this.contentArray[num]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
