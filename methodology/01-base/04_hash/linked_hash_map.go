package main

import "fmt"

// MyLinkedHashMap 链表加强哈希表
type MyLinkedHashMap struct {
	head *Node
	tail *Node
	m    map[string]*Node
}

type Node struct {
	key  string
	val  int
	next *Node
	prev *Node
}

func Constructor() MyLinkedHashMap {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return MyLinkedHashMap{
		head: head,
		tail: tail,
		m:    make(map[string]*Node),
	}
}

func (this *MyLinkedHashMap) Get(key string) int {
	if node, ok := this.m[key]; ok {
		return node.val
	}
	return -1
}

func (this *MyLinkedHashMap) Put(key string, val int) {
	// 若为新插入的节点，则同时插入链表和 map
	if _, ok := this.m[key]; !ok {
		// 插入新的 Node
		node := &Node{key: key, val: val}
		this.addLastNode(node)
		this.m[key] = node
		return
	}
	// 若存在，则替换之前的 val
	this.m[key].val = val
}

func (this *MyLinkedHashMap) Remove(key string) {
	// 若 key 本不存在，直接返回
	if _, ok := this.m[key]; !ok {
		return
	}
	// 若 key 存在，则需要同时在哈希表和链表中删除
	node := this.m[key]
	delete(this.m, key)
	this.removeNode(node)
}

func (this *MyLinkedHashMap) ContainsKey(key string) bool {
	_, ok := this.m[key]
	return ok
}

func (this *MyLinkedHashMap) Keys() []string {
	keyList := make([]string, 0)
	for p := this.head.next; p != this.tail; p = p.next {
		keyList = append(keyList, p.key)
	}
	return keyList
}

func (this *MyLinkedHashMap) addLastNode(x *Node) {
	temp := this.tail.prev
	// temp <-> tail

	x.next = this.tail
	x.prev = temp
	// temp <- x -> tail

	temp.next = x
	this.tail.prev = x
	// temp <-> x <-> tail
}

func (this *MyLinkedHashMap) removeNode(x *Node) {
	prev := x.prev
	next := x.next
	// prev <-> x <-> next

	prev.next = next
	next.prev = prev

	x.next = nil
	x.prev = nil
}

func main() {
	myMap := Constructor()
	myMap.Put("a", 1)
	myMap.Put("b", 2)
	myMap.Put("c", 3)
	myMap.Put("d", 4)
	myMap.Put("e", 5)

	// output: a b c d e
	fmt.Println(myMap.Keys())

	myMap.Remove("c")

	// output: a b d e
	fmt.Println(myMap.Keys())
}
