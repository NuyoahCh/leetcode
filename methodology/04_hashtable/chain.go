package main

import (
	"container/list"
	"errors"
	"fmt"
)

type MyChainingHashMap[K comparable, V any] struct {
	// 哈希表的底层数组，每个数组元素是一个链表，链表中每个节点是 KVNode 存储键值对
	table []*list.List

	// 哈希表中存入的键值对个数
	size int
}

// KVNode 拉链法使用的单链表节点，存储 key-value 对
type KVNode[K comparable, V any] struct {
	key   K
	value V
}

// INIT_CAP 底层数组的初始容量
const INIT_CAP = 4

// NewMyChainingHashMap 创建一个新的 MyChainingHashMap，使用默认的初始容量
func NewMyChainingHashMap[K comparable, V any]() *MyChainingHashMap[K, V] {
	return NewMyChainingHashMapWithCapacity[K, V](INIT_CAP)
}

// NewMyChainingHashMapWithCapacity 使用指定的初始容量创建 MyChainingHashMap
func NewMyChainingHashMapWithCapacity[K comparable, V any](initCapacity int) *MyChainingHashMap[K, V] {
	if initCapacity < 1 {
		initCapacity = 1
	}
	m := &MyChainingHashMap[K, V]{
		table: make([]*list.List, initCapacity),
	}
	for i := range m.table {
		m.table[i] = list.New()
	}
	return m
}

// Put 添加 key -> val 键值对
func (m *MyChainingHashMap[K, V]) Put(key K, val V) error {
	if key == *new(K) {
		// 相当于 nil 检查
		return errors.New("key is null")
	}
	index := m.hash(key)
	l := m.table[index]
	// 如果 key 之前存在，则修改对应的 val
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode[K, V])
		if node.key == key {
			e.Value = KVNode[K, V]{key, val}
			return nil
		}
	}
	// 如果 key 之前不存在，则插入，size 增加
	l.PushBack(KVNode[K, V]{key, val})
	m.size++

	// 如果元素数量超过了负载因子，进行扩容
	if m.size >= len(m.table)*3/4 {
		m.resize(len(m.table) * 2)
	}
	return nil
}

// Remove 删除 key 和对应的 val
func (m *MyChainingHashMap[K, V]) Remove(key K) error {
	if key == *new(K) {
		return errors.New("key is null")
	}
	index := m.hash(key)
	l := m.table[index]
	// 如果 key 存在，则删除，size 减少
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode[K, V])
		if node.key == key {
			l.Remove(e)
			m.size--

			// 缩容，当负载因子小于 0.125 时，缩容
			if m.size <= len(m.table)/8 {
				m.resize(len(m.table) / 2)
			}
			return nil
		}
	}
	return nil
}

// Get 返回 key 对应的 val，如果 key 不存在，则返回 nil
func (m *MyChainingHashMap[K, V]) Get(key K) *V {
	if key == *new(K) {
		return nil
	}
	index := m.hash(key)
	l := m.table[index]
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode[K, V])
		if node.key == key {
			return &node.value
		}
	}
	return nil
}

// 返回所有 key
func (m *MyChainingHashMap[K, V]) Keys() []K {
	var keys []K
	for _, l := range m.table {
		for e := l.Front(); e != nil; e = e.Next() {
			node := e.Value.(KVNode[K, V])
			keys = append(keys, node.key)
		}
	}
	return keys
}

// 返回当前哈希表中的键值对个数
func (m *MyChainingHashMap[K, V]) Size() int {
	return m.size
}

// 哈希函数，将键映射到 table 的索引
func (m *MyChainingHashMap[K, V]) hash(key K) int {
	return int(fmt.Sprintf("%v", key)[0]) % len(m.table)
}

// 扩展或缩小哈希表的容量
func (m *MyChainingHashMap[K, V]) resize(newCap int) {
	if newCap < 1 {
		newCap = 1
	}
	newMap := NewMyChainingHashMapWithCapacity[K, V](newCap)
	for _, l := range m.table {
		for e := l.Front(); e != nil; e = e.Next() {
			node := e.Value.(KVNode[K, V])
			newMap.Put(node.key, node.value)
		}
	}
	m.table = newMap.table
}

// 测试代码
func main() {
	map1 := NewMyChainingHashMap[int, int]()
	map1.Put(1, 1)
	map1.Put(2, 2)
	map1.Put(3, 3)
	fmt.Println(*map1.Get(1)) // 1
	fmt.Println(*map1.Get(2)) // 2

	map1.Put(1, 100)
	fmt.Println(*map1.Get(1)) // 100

	map1.Remove(2)
	fmt.Println(map1.Get(2)) // nil

	fmt.Println(map1.Keys()) // [1, 3]

	map1.Remove(1)
	map1.Remove(3)
	fmt.Println(map1.Get(1)) // nil
}
