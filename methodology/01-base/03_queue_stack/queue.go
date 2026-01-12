package main

import (
	"errors"
	"fmt"
)

// Node 队列节点
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Queue 基于链表的队列
type Queue[T any] struct {
	head *Node[T] // 队头：取数据
	tail *Node[T] // 队尾：存数据
	size int
}

// Enqueue 入队：O(1)
func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{Value: value}

	if q.tail == nil {
		// 如果队列为空，头尾都指向新节点
		q.head = newNode
		q.tail = newNode
	} else {
		// 1. 当前尾节点的 Next 指向新节点
		q.tail.Next = newNode
		// 2. 更新尾指针指向新节点
		q.tail = newNode
	}
	q.size++
}

// Dequeue 出队：O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.head == nil {
		return zero, errors.New("queue is empty")
	}

	// 1. 获取头部数据
	result := q.head.Value

	// 2. 头指针后移
	q.head = q.head.Next

	// 3. 边界处理：如果队列变空了，tail 也要置空
	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return result, nil
}

func (q *Queue[T]) Size() int {
	return q.size
}

func main() {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println("Queue Size:", q.Size()) // 3

	val1, _ := q.Dequeue()
	fmt.Println("Dequeued:", val1) // 1 (先进先出)

	val2, _ := q.Dequeue()
	fmt.Println("Dequeued:", val2) // 2
}
