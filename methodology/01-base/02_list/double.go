package main

import "fmt"

// DoubleNode 双链表节点
type DoubleNode[T any] struct {
	Value T
	Prev  *DoubleNode[T] // 指向前一个
	Next  *DoubleNode[T] // 指向后一个
}

// DoubleLinkedList 双链表容器
type DoubleLinkedList[T any] struct {
	Head *DoubleNode[T]
	Tail *DoubleNode[T]
}

// AddToTail 尾部追加：O(1)
// 双链表通常维护一个 Tail 指针，使得尾插非常快
func (list *DoubleLinkedList[T]) AddToTail(val T) {
	newNode := &DoubleNode[T]{Value: val}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	// 建立双向连接
	list.Tail.Next = newNode // 旧尾指向新节点
	newNode.Prev = list.Tail // 新节点指向旧尾
	list.Tail = newNode      // 更新尾指针
}

// RemoveNode 删除指定节点：O(1)
// 前提：你已经拿到了要删除的这个 node 对象
func (list *DoubleLinkedList[T]) RemoveNode(node *DoubleNode[T]) {
	if node == nil {
		return
	}

	// 1. 处理前驱
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// 如果没有前驱，说明删除的是头节点
		list.Head = node.Next
	}

	// 2. 处理后继
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		// 如果没有后继，说明删除的是尾节点
		list.Tail = node.Prev
	}

	// 3. 清理（可选，帮助GC）
	node.Prev = nil
	node.Next = nil
}

// 打印链表（正向）
func (list *DoubleLinkedList[T]) String() string {
	s := "Head <-> "
	current := list.Head
	for current != nil {
		s += fmt.Sprintf("%v <-> ", current.Value)
		current = current.Next
	}
	return s + "nil"
}

func main() {
	dl := DoubleLinkedList[string]{}
	dl.AddToTail("A")
	dl.AddToTail("B")
	dl.AddToTail("C")

	fmt.Println("初始:", dl.String()) // A <-> B <-> C

	// 模拟场景：我们已经定位到了中间的节点 B
	nodeB := dl.Head.Next

	// 执行 O(1) 删除
	dl.RemoveNode(nodeB)

	fmt.Println("删除B后:", dl.String()) // A <-> C
}
