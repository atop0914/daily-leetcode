package main

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */

// @lc code=start
type RandomizedCollection struct {
	// 存放元素val及对应的位置合集
	contentSet   map[int]map[int]struct{}
	contentArray []int
}

func Constructor_381() RandomizedCollection {
	return RandomizedCollection{
		contentSet:   make(map[int]map[int]struct{}),
		contentArray: make([]int, 0),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	this.contentArray = append(this.contentArray, val)

	var exist bool

	if _, ok := this.contentSet[val]; ok {
		this.contentSet[val][len(this.contentArray)-1] = struct{}{}
	} else {
		exist = true
		set := make(map[int]struct{}, 0)
		set[len(this.contentArray)-1] = struct{}{}
		this.contentSet[val] = set
	}

	return exist

}

func (this *RandomizedCollection) Remove(val int) bool {
	var i int // 待交换的数组下标

	_, exist := this.contentSet[val]

	// 不存在
	if !exist {
		return false
	}

	// 数组中只有一个元素
	if len(this.contentSet) == 1 && len(this.contentArray) == 1 {
		this.contentSet = make(map[int]map[int]struct{})
		this.contentArray = make([]int, 0)
		return true
	}

	for indexs := range this.contentSet[val] {
		i = indexs
		break
	}

	// 更改map中的索引数值
	delete(this.contentSet[val], i)
	if len(this.contentSet[val]) == 0 {
		delete(this.contentSet, val)
	}
	lastItem := this.contentArray[len(this.contentArray)-1]
	delete(this.contentSet[lastItem], len(this.contentArray)-1)

	// 是否存在 如果要删除的元素刚好是最后一个那不用额外对她进行指定
	if i != len(this.contentArray)-1 { // ***非常重要***因为当删除元素刚好是队尾元素的时候 不需要去修改它原先的索引值了
		if _, ok := this.contentSet[lastItem]; ok {
			this.contentSet[lastItem][i] = struct{}{}
		} else {
			l := make(map[int]struct{})
			l[i] = struct{}{}
			this.contentSet[lastItem] = l
		}
	}

	// 将待交换的元素与数组最后一个元素交换
	this.contentArray[i], this.contentArray[len(this.contentArray)-1] = this.contentArray[len(this.contentArray)-1], this.contentArray[i]

	this.contentArray = this.contentArray[:len(this.contentArray)-1]

	return true
}

func (this *RandomizedCollection) GetRandom() int {
	if len(this.contentArray) == 0 {
		return 0
	}
	return this.contentArray[rand.Intn(len(this.contentArray))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

// func main() {
// 	c := Constructor()
// 	fmt.Println(c.Insert(10))
// 	fmt.Println(c.Insert(10))
// 	fmt.Println(c.Insert(20))
// 	fmt.Println(c.Insert(20))
// 	fmt.Println(c.Insert(30))
// 	fmt.Println(c.Insert(30))
// 	fmt.Println(c.Remove(10))
// 	fmt.Println(c.Remove(20))
// 	fmt.Println(c.Remove(20))
// 	fmt.Println(c.Remove(10))
// 	fmt.Println(c.Remove(30))
// 	fmt.Println(c.Insert(40))
// 	fmt.Println(c.Remove(30))
// 	fmt.Println(c.Remove(30))
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// 	fmt.Println(c.GetRandom())
// }
