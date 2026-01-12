package main

import (
	"errors"
	"fmt"
)

type MyPriorityQueue struct {
	// 堆数组
	heap []interface{}

	// 堆中元素的数量
	size int

	// 元素比较器
	comparator func(x, y interface{}) int
}

// 构造函数
func NewMyPriorityQueue(capacity int, comparator func(x, y interface{}) int) *MyPriorityQueue {
	return &MyPriorityQueue{
		heap:       make([]interface{}, capacity),
		size:       0,
		comparator: comparator,
	}
}

// 返回堆的大小
func (pq *MyPriorityQueue) Size() int {
	return pq.size
}

// 判断堆是否为空
func (pq *MyPriorityQueue) IsEmpty() bool {
	return pq.size == 0
}

// 父节点的索引
func (pq *MyPriorityQueue) Parent(node int) int {
	return (node - 1) / 2
}

// 左子节点的索引
func (pq *MyPriorityQueue) Left(node int) int {
	return node*2 + 1
}

// 右子节点的索引
func (pq *MyPriorityQueue) Right(node int) int {
	return node*2 + 2
}

// 交换数组的两个元素
func (pq *MyPriorityQueue) Swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

// 查，返回堆顶元素，时间复杂度 O(1)
func (pq *MyPriorityQueue) Peek() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	return pq.heap[0], nil
}

// 增，向堆中插入一个元素，时间复杂度 O(logN)
func (pq *MyPriorityQueue) Push(x interface{}) {
	// 扩容
	if pq.size == len(pq.heap) {
		pq.resize(2 * len(pq.heap))
	}
	// 把新元素追加到最后
	pq.heap[pq.size] = x
	// 然后上浮到正确位置
	pq.swim(pq.size)
	pq.size++
}

// 删，删除堆顶元素，时间复杂度 O(logN)
func (pq *MyPriorityQueue) Pop() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	res := pq.heap[0]
	// 把堆底元素放到堆顶
	pq.Swap(0, pq.size-1)
	// 避免对象游离
	pq.heap[pq.size-1] = nil
	pq.size--
	// 然后下沉到正确位置
	pq.sink(0)
	// 缩容
	if pq.size > 0 && pq.size == len(pq.heap)/4 {
		pq.resize(len(pq.heap) / 2)
	}
	return res, nil
}

// 上浮操作，时间复杂度是树高 O(logN)
func (pq *MyPriorityQueue) swim(node int) {
	for node > 0 && pq.comparator(pq.heap[pq.Parent(node)], pq.heap[node]) > 0 {
		pq.Swap(pq.Parent(node), node)
		node = pq.Parent(node)
	}
}

// 下沉操作，时间复杂度是树高 O(logN)
func (pq *MyPriorityQueue) sink(node int) {
	for pq.Left(node) < pq.size {
		// 比较自己和左右子节点，看看谁最小
		minNode := node
		if pq.Left(node) < pq.size && pq.comparator(pq.heap[pq.Left(node)], pq.heap[minNode]) < 0 {
			minNode = pq.Left(node)
		}
		if pq.Right(node) < pq.size && pq.comparator(pq.heap[pq.Right(node)], pq.heap[minNode]) < 0 {
			minNode = pq.Right(node)
		}
		if minNode == node {
			break
		}
		// 如果左右子节点中有比自己小的，就交换
		pq.Swap(node, minNode)
		node = minNode
	}
}

// 调整堆的大小
func (pq *MyPriorityQueue) resize(capacity int) {
	newHeap := make([]interface{}, capacity)
	for i := 0; i < pq.size; i++ {
		newHeap[i] = pq.heap[i]
	}
	pq.heap = newHeap
}

func main() {
	pq := NewMyPriorityQueue(3, func(x, y interface{}) int {
		a := x.(int)
		b := y.(int)
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	pq.Push(3)
	pq.Push(1)
	pq.Push(4)
	pq.Push(1)
	pq.Push(5)
	pq.Push(9)

	// 1 1 3 4 5 9
	for !pq.IsEmpty() {
		item, _ := pq.Pop()
		fmt.Println(item)
	}
}
