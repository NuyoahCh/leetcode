package main

import (
	"container/list"
	"fmt"
)

type MyListDeque struct {
	list *list.List
}

func NewMyListDeque() *MyListDeque {
	return &MyListDeque{list: list.New()}
}

// AddFirst 从队头插入元素，时间复杂度 O(1)
func (d *MyListDeque) AddFirst(e interface{}) {
	d.list.PushFront(e)
}

// AddLast 从队尾插入元素，时间复杂度 O(1)
func (d *MyListDeque) AddLast(e interface{}) {
	d.list.PushBack(e)
}

// RemoveFirst 从队头删除元素，时间复杂度 O(1)
func (d *MyListDeque) RemoveFirst() interface{} {
	if elem := d.list.Front(); elem != nil {
		return d.list.Remove(elem)
	}
	return nil
}

// RemoveLast 从队尾删除元素，时间复杂度 O(1)
func (d *MyListDeque) RemoveLast() interface{} {
	if elem := d.list.Back(); elem != nil {
		return d.list.Remove(elem)
	}
	return nil
}

// PeekFirst 查看队头元素，时间复杂度 O(1)
func (d *MyListDeque) PeekFirst() interface{} {
	if elem := d.list.Front(); elem != nil {
		return elem.Value
	}
	return nil
}

// PeekLast 查看队尾元素，时间复杂度 O(1)
func (d *MyListDeque) PeekLast() interface{} {
	if elem := d.list.Back(); elem != nil {
		return elem.Value
	}
	return nil
}

func main() {
	deque := NewMyListDeque()
	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddLast(3)
	deque.AddLast(4)

	fmt.Println(deque.RemoveFirst()) // 2
	fmt.Println(deque.RemoveLast())  // 4
	fmt.Println(deque.PeekFirst())   // 1
	fmt.Println(deque.PeekLast())    // 3
}
