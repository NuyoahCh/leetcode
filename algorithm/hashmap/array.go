package main

import (
	"fmt"
	"math/rand"
)

// ArrayNode 定义数组哈希
type ArrayNode struct {
	key int
	val int
}

// MyArrayHashMap 初始化数组哈希
type MyArrayHashMap struct {
	// 存储 key 和 key 在 arr 中的索引
	m map[int]int
	// 真正存储 key-value 的数组
	arr []ArrayNode
}

func NewMyArrayHashMap() *MyArrayHashMap {
	return &MyArrayHashMap{
		m:   make(map[int]int),
		arr: make([]ArrayNode, 0),
	}
}

func (m *MyArrayHashMap) Get(key int) int {
	if _, ok := m.m[key]; !ok {
		return -1
	}
	return m.arr[m.m[key]].val
}

func (m *MyArrayHashMap) Put(key int, val int) {
	if m.containsKey(key) {
		// 修改
		i := m.m[key]
		m.arr[i].val = val
		return
	}

	// 新增
	m.arr = append(m.arr, ArrayNode{key, val})
	m.m[key] = len(m.arr) - 1
}

func (m *MyArrayHashMap) Remove(key int) {
	if _, ok := m.m[key]; !ok {
		return
	}
	index := m.m[key]
	node := m.arr[index]

	// 1. 最后一个元素 e 和第 index 个元素 node 换位置
	e := m.arr[len(m.arr)-1]
	m.arr[index] = e
	m.arr[len(m.arr)-1] = node

	// 2. 修改 map 中 e.key 对应的索引
	m.m[e.key] = index

	// 3. 在数组中删除最后一个元素
	m.arr = m.arr[:len(m.arr)-1]

	// 4. 在 map 中删除 node.key
	delete(m.m, node.key)
}

// 随机弹出一个键
func (m *MyArrayHashMap) randomKey() int {
	n := len(m.arr)
	randomIndex := rand.Intn(n)
	return m.arr[randomIndex].key
}

func (m *MyArrayHashMap) containsKey(key int) bool {
	_, ok := m.m[key]
	return ok
}

func (m *MyArrayHashMap) size() int {
	return len(m.m)
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
