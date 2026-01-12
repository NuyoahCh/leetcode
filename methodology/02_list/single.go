package main

import "fmt"

// SingleNode 单链表节点
type SingleNode[T any] struct {
	Value T
	Next  *SingleNode[T]
}

// SingleLinkedList 单链表容器
type SingleLinkedList[T any] struct {
	Head *SingleNode[T]
	Size int
}

// Get 查找：O(N)
// 需要从头遍历
func (list *SingleLinkedList[T]) Get(index int) (T, bool) {
	var zero T
	if index < 0 || index >= list.Size {
		return zero, false
	}

	current := list.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, true
}

// AddToHead 头部插入：O(1)
// 这是单链表最快的操作
func (list *SingleLinkedList[T]) AddToHead(val T) {
	newNode := &SingleNode[T]{Value: val}
	newNode.Next = list.Head // 新节点指向旧头
	list.Head = newNode      // 头指针指向新节点
	list.Size++
}

// Update 修改：O(N)
func (list *SingleLinkedList[T]) Update(index int, val T) {
	current := list.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}
	if current != nil {
		current.Value = val
	}
}

// 打印链表
func (list *SingleLinkedList[T]) String() string {
	s := ""
	current := list.Head
	for current != nil {
		s += fmt.Sprintf("%v -> ", current.Value)
		current = current.Next
	}
	return s + "nil"
}

func main() {
	sl := SingleLinkedList[int]{}
	sl.AddToHead(10)
	sl.AddToHead(20)                 // 20 变成新的头
	fmt.Println("单链表:", sl.String()) // 20 -> 10 -> nil
}
