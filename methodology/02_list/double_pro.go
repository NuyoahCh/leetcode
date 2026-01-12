package main

import (
	"errors"
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
	prev *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// NewMyLinkedList 虚拟头尾节点
func NewMyLinkedList() *MyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &MyLinkedList{head: head, tail: tail, size: 0}
}

// ***** 增 *****

func (list *MyLinkedList) AddLast(e interface{}) {
	x := &Node{val: e}
	temp := list.tail.prev
	// temp <-> x
	temp.next = x
	x.prev = temp

	x.next = list.tail
	list.tail.prev = x
	// temp <-> x <-> tail
	list.size++
}

func (list *MyLinkedList) AddFirst(e interface{}) {
	x := &Node{val: e}
	temp := list.head.next
	// head <-> temp
	temp.prev = x
	x.next = temp

	list.head.next = x
	x.prev = list.head
	// head <-> x <-> temp
	list.size++
}

func (list *MyLinkedList) Add(index int, element interface{}) error {
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}
	if index == list.size {
		list.AddLast(element)
		return nil
	}

	// 找到 index 对应的 Node
	p := list.getNode(index)
	temp := p.prev
	// temp <-> p

	// 新要插入的 Node
	x := &Node{val: element}

	p.prev = x
	temp.next = x

	x.prev = temp
	x.next = p

	// temp <-> x <-> p

	list.size++
	return nil
}

// ***** 删 *****

func (list *MyLinkedList) RemoveFirst() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements to remove")
	}
	// 虚拟节点的存在是我们不用考虑空指针的问题
	x := list.head.next
	temp := x.next
	// head <-> x <-> temp
	list.head.next = temp
	temp.prev = list.head

	// head <-> temp

	list.size--
	return x.val, nil
}

func (list *MyLinkedList) RemoveLast() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements to remove")
	}
	x := list.tail.prev
	temp := x.prev
	// temp <-> x <-> tail

	list.tail.prev = temp
	temp.next = list.tail

	// temp <-> tail

	list.size--
	return x.val, nil
}

func (list *MyLinkedList) Remove(index int) (interface{}, error) {
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 找到 index 对应的 Node
	x := list.getNode(index)
	prev := x.prev
	next := x.next
	// prev <-> x <-> next
	prev.next = next
	next.prev = prev

	list.size--

	return x.val, nil
}

// ***** 查 *****

func (list *MyLinkedList) Get(index int) (interface{}, error) {
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 找到 index 对应的 Node
	p := list.getNode(index)

	return p.val, nil
}

func (list *MyLinkedList) GetFirst() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements in the list")
	}

	return list.head.next.val, nil
}

func (list *MyLinkedList) GetLast() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements in the list")
	}

	return list.tail.prev.val, nil
}

// ***** 改 *****

func (list *MyLinkedList) Set(index int, val interface{}) (interface{}, error) {
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 找到 index 对应的 Node
	p := list.getNode(index)

	oldVal := p.val
	p.val = val

	return oldVal, nil
}

// ***** 其他工具函数 *****

func (list *MyLinkedList) Size() int {
	return list.size
}

func (list *MyLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *MyLinkedList) getNode(index int) *Node {
	p := list.head.next
	// TODO: 可以优化，通过 index 判断从 head 还是 tail 开始遍历
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}

func (list *MyLinkedList) isElementIndex(index int) bool {
	return index >= 0 && index < list.size
}

func (list *MyLinkedList) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}

// 检查 index 索引位置是否可以存在元素
func (list *MyLinkedList) checkElementIndex(index int) error {
	if !list.isElementIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

// 检查 index 索引位置是否可以添加元素
func (list *MyLinkedList) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *MyLinkedList) Display() {
	fmt.Printf("size = %d\n", list.size)
	p := list.head.next
	for p != list.tail {
		fmt.Printf("%v <-> ", p.val)
		p = p.next
	}
	fmt.Println("null\n")
}

func main() {
	list := NewMyLinkedList()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddFirst(0)
	list.Add(2, 100)

	list.Display()
	// size = 5
	// 0 <-> 1 <-> 100 <-> 2 <-> 3 <-> null
}
