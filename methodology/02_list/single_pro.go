package main

import (
	"errors"
)

// Node01 节点结构
type Node01[E any] struct {
	val  E
	next *Node01[E]
}

// MyLinkedList2 链表结构
type MyLinkedList2[E any] struct {
	head *Node01[E]
	// 实际的尾部节点引用
	tail  *Node01[E]
	size_ int
}

// NewMyLinkedList2 创建一个新的链表
func NewMyLinkedList2[E any]() *MyLinkedList2[E] {
	head := &Node01[E]{}
	return &MyLinkedList2[E]{head: head, tail: head, size_: 0}
}

// AddFirst 在头部添加元素
func (list *MyLinkedList2[E]) AddFirst(e E) {
	newNode01 := &Node01[E]{val: e}
	newNode01.next = list.head.next
	list.head.next = newNode01
	if list.size_ == 0 {
		list.tail = newNode01
	}
	list.size_++
}

// AddLast 在尾部添加元素
func (list *MyLinkedList2[E]) AddLast(e E) {
	newNode01 := &Node01[E]{val: e}
	list.tail.next = newNode01
	list.tail = newNode01
	list.size_++
}

// Add 在指定索引处添加元素
func (list *MyLinkedList2[E]) Add(index int, element E) error {
	if index < 0 || index > list.size_ {
		return errors.New("index out of bounds")
	}

	if index == list.size_ {
		list.AddLast(element)
		return nil
	}

	prev := list.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	newNode01 := &Node01[E]{val: element}
	newNode01.next = prev.next
	prev.next = newNode01
	list.size_++
	return nil
}

// RemoveFirst 移除头部元素
func (list *MyLinkedList2[E]) RemoveFirst() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements to remove")
	}
	first := list.head.next
	list.head.next = first.next
	if list.size_ == 1 {
		list.tail = list.head
	}
	list.size_--
	return first.val, nil
}

// RemoveLast 移除尾部元素
func (list *MyLinkedList2[E]) RemoveLast() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements to remove")
	}

	prev := list.head
	for prev.next != list.tail {
		prev = prev.next
	}
	val := list.tail.val
	prev.next = nil
	list.tail = prev
	list.size_--
	return val, nil
}

// Remove 在指定索引处移除元素
func (list *MyLinkedList2[E]) Remove(index int) (E, error) {
	if index < 0 || index >= list.size_ {
		return *new(E), errors.New("index out of bounds")
	}

	prev := list.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	Node01ToRemove := prev.next
	prev.next = Node01ToRemove.next
	// 删除的是最后一个元素
	if index == list.size_-1 {
		list.tail = prev
	}
	list.size_--
	return Node01ToRemove.val, nil
}

// GetFirst 获取头部元素
func (list *MyLinkedList2[E]) GetFirst() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements in the list")
	}
	return list.head.next.val, nil
}

// GetLast 获取尾部元素
func (list *MyLinkedList2[E]) GetLast() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements in the list")
	}
	return list.tail.val, nil
}

// Get 获取指定索引的元素
func (list *MyLinkedList2[E]) Get(index int) (E, error) {
	if index < 0 || index >= list.size_ {
		return *new(E), errors.New("index out of bounds")
	}
	return list.getNode01(index).val, nil
}

// Set 更新指定索引的元素
func (list *MyLinkedList2[E]) Set(index int, element E) (E, error) {
	if index < 0 || index >= list.size_ {
		return *new(E), errors.New("index out of bounds")
	}
	Node01 := list.getNode01(index)
	oldVal := Node01.val
	Node01.val = element
	return oldVal, nil
}

// Size 获取链表大小
func (list *MyLinkedList2[E]) Size() int {
	return list.size_
}

// IsEmpty 检查链表是否为空
func (list *MyLinkedList2[E]) IsEmpty() bool {
	return list.size_ == 0
}

// getNode01 返回指定索引的节点
func (list *MyLinkedList2[E]) getNode01(index int) *Node01[E] {
	p := list.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}
