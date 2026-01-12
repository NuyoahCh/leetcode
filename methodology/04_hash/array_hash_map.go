package main

import (
	"fmt"
	"math/rand"
)

type Node01 struct {
	key int
	val int
}

type MyArrayHashMap struct {
	// 存储 key 和 key 在 arr 中的索引
	m map[int]int
	// 真正存储 key-value 的数组
	arr []Node01
}

func NewMyArrayHashMap() *MyArrayHashMap {
	return &MyArrayHashMap{
		m:   make(map[int]int),
		arr: make([]Node01, 0),
	}
}

func (this *MyArrayHashMap) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	return this.arr[this.m[key]].val
}

func (this *MyArrayHashMap) Put(key int, val int) {
	if this.containsKey(key) {
		// 修改
		i := this.m[key]
		this.arr[i].val = val
		return
	}

	// 新增
	this.arr = append(this.arr, Node01{key, val})
	this.m[key] = len(this.arr) - 1
}

func (this *MyArrayHashMap) Remove(key int) {
	if _, ok := this.m[key]; !ok {
		return
	}
	index := this.m[key]
	Node01 := this.arr[index]

	// 1. 最后一个元素 e 和第 index 个元素 Node01 换位置
	e := this.arr[len(this.arr)-1]
	this.arr[index] = e
	this.arr[len(this.arr)-1] = Node01

	// 2. 修改 map 中 e.key 对应的索引
	this.m[e.key] = index

	// 3. 在数组中删除最后一个元素
	this.arr = this.arr[:len(this.arr)-1]

	// 4. 在 map 中删除 Node01.key
	delete(this.m, Node01.key)
}

// 随机弹出一个键
func (this *MyArrayHashMap) randomKey() int {
	n := len(this.arr)
	randomIndex := rand.Intn(n)
	return this.arr[randomIndex].key
}

func (this *MyArrayHashMap) containsKey(key int) bool {
	_, ok := this.m[key]
	return ok
}

func (this *MyArrayHashMap) size() int {
	return len(this.m)
}

func main() {
	arrMap := NewMyArrayHashMap()
	arrMap.Put(1, 1)
	arrMap.Put(2, 2)
	arrMap.Put(3, 3)
	arrMap.Put(4, 4)
	arrMap.Put(5, 5)

	fmt.Println(arrMap.Get(1)) // 1
	fmt.Println(arrMap.randomKey())

	arrMap.Remove(4)
	fmt.Println(arrMap.randomKey())
	fmt.Println(arrMap.randomKey())
}
