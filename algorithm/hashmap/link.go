package main

import "fmt"

// Node 链表的节点
type Node struct {
	key  string
	val  int
	next *Node
	prev *Node
}

// MyLinkedHashMap 初始化链表加强哈希表
type MyLinkedHashMap struct {
	head *Node
	tail *Node
	m    map[string]*Node
}

// Constructor 构造器
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

// Get 获取 node 节点值
func (m *MyLinkedHashMap) Get(key string) int {
	// key value 形式查找
	if node, ok := m.m[key]; ok {
		return node.val
	}
	return -1
}

// Put 写入 node 节点值
func (m *MyLinkedHashMap) Put(key string, val int) {
	// 若为新插入的节点，则同时插入链表和 map
	if _, ok := m.m[key]; !ok {
		// 插入新的 Node
		node := &Node{key: key, val: val}
		m.addLastNode(node)
		m.m[key] = node
		return
	}
	// 若存在，则替换之前的 val
	m.m[key].val = val
}

// Remove 删除 node 节点值
func (m *MyLinkedHashMap) Remove(key string) {
	// 若 key 本不存在，直接返回
	if _, ok := m.m[key]; !ok {
		return
	}
	// 若 key 存在，则需要同时在哈希表和链表中删除
	node := m.m[key]
	delete(m.m, key)
	m.removeNode(node)
}

// ContainsKey 是否存在 key 节点
func (m *MyLinkedHashMap) ContainsKey(key string) bool {
	_, ok := m.m[key]
	return ok
}

// Keys 将 key 添加到 list 中
func (m *MyLinkedHashMap) Keys() []string {
	keyList := make([]string, 0)
	for p := m.head.next; p != m.tail; p = p.next {
		keyList = append(keyList, p.key)
	}
	return keyList
}

// 在尾部添加 node 节点
func (m *MyLinkedHashMap) addLastNode(x *Node) {
	temp := m.tail.prev
	// temp <-> tail

	x.next = m.tail
	x.prev = temp

	temp.next = x
	m.tail.prev = x
}

// 删除 node 节点
func (m *MyLinkedHashMap) removeNode(x *Node) {
	prev := x.prev
	next := x.next
	// prev <-> x <-> next

	prev.next = next
	next.prev = prev

	// 置 nil 是一个好习惯
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
