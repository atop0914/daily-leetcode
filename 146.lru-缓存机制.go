package main

import (
	"fmt"
	"math"
	"time"
)

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	cap int
	m   map[int]int   // 存放 k v 数据
	hot map[int]int64 // 存放 k times 数据
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		cap: capacity,
		m:   make(map[int]int, capacity),
		hot: make(map[int]int64, capacity),
	}
}

func (this *LRUCache) Get(key int) int {

	if value, ok := this.m[key]; ok {
		this.hot[key] = time.Now().Unix()
		return value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.m[key]; ok {
		this.m[key] = value
	} else {
		if len(this.m) == this.cap { // 需要扩容
			min := int64(math.MaxInt64)
			kmin := -1
			for k, v := range this.hot {
				if v < min {
					min = v
					kmin = k
				}
			}
			delete(this.m, kmin)
			delete(this.hot, kmin)
		}
		this.m[key] = value
	}
	this.hot[key] = time.Now().Unix()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
