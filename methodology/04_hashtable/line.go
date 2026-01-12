package main

import (
	"fmt"
	"hash/fnv"
)

type KVNode02[K comparable, V any] struct {
	key K
	val V
}

// MyLinearProbingHashMap1 线性探查哈希表
type MyLinearProbingHashMap1[K comparable, V any] struct {
	table []*KVNode02[K, V]
	size  int
}

func NewMyLinearProbingHashMap1[K comparable, V any]() *MyLinearProbingHashMap1[K, V] {
	return NewMyLinearProbingHashMap1WithCapacity[K, V](INIT_CAP)
}

func NewMyLinearProbingHashMap1WithCapacity[K comparable, V any](initCapacity int) *MyLinearProbingHashMap1[K, V] {
	return &MyLinearProbingHashMap1[K, V]{
		table: make([]*KVNode02[K, V], initCapacity),
	}
}

// **** 增/改 ****
func (m *MyLinearProbingHashMap1[K, V]) Put(key K, val V) error {
	// 负载因子默认设为 0.75，超过则扩容
	if m.size >= len(m.table)*3/4 {
		m.resize(len(m.table) * 2)
	}

	index := m.getKeyIndex(key)
	// key 已存在，修改对应的 val
	if m.table[index] != nil && m.table[index].key == key {
		m.table[index].val = val
		return nil
	}

	// key 不存在，在空位插入
	m.table[index] = &KVNode02[K, V]{key, val}
	m.size++
	return nil
}

// **** 删 ****
func (m *MyLinearProbingHashMap1[K, V]) Remove(key K) error {
	// 缩容，当负载因子小于 0.125 时，缩容
	if m.size <= len(m.table)/8 {
		m.resize(len(m.table) / 2)
	}

	index := m.getKeyIndex(key)

	if m.table[index] == nil {
		// key 不存在，不需要 remove
		return nil
	}

	// 开始 remove
	m.table[index] = nil
	m.size--
	// 保持元素连续性，进行 rehash
	index = (index + 1) % len(m.table)
	for m.table[index] != nil {
		entry := m.table[index]
		m.table[index] = nil
		m.size--
		m.Put(entry.key, entry.val)
		index = (index + 1) % len(m.table)
	}
	return nil
}

// **** 查 ****
func (m *MyLinearProbingHashMap1[K, V]) Get(key K) (V, error) {
	index := m.getKeyIndex(key)
	if m.table[index] == nil {
		return *new(V), nil
	}
	return m.table[index].val, nil
}

func (m *MyLinearProbingHashMap1[K, V]) ContainsKey(key K) bool {
	index := m.getKeyIndex(key)
	return m.table[index] != nil
}

// 返回所有 key（顺序不固定）
func (m *MyLinearProbingHashMap1[K, V]) Keys() []K {
	keys := []K{}
	for _, entry := range m.table {
		if entry != nil {
			keys = append(keys, entry.key)
		}
	}
	return keys
}

// **** 其他工具函数 ****
func (m *MyLinearProbingHashMap1[K, V]) Size() int {
	return m.size
}

// 哈希函数，将键映射到 table 的索引
func (m *MyLinearProbingHashMap1[K, V]) hash(key K) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) & 0x7fffffff % len(m.table)
}

// 对 key 进行线性探查，返回一个索引
// 如果 key 不存在，返回的就是下一个为 null 的索引，可用于插入
func (m *MyLinearProbingHashMap1[K, V]) getKeyIndex(key K) int {
	index := m.hash(key)
	for m.table[index] != nil {
		if m.table[index].key == key {
			return index
		}
		index = (index + 1) % len(m.table)
	}
	return index
}

func (m *MyLinearProbingHashMap1[K, V]) resize(newCap int) {
	newMap := NewMyLinearProbingHashMap1WithCapacity[K, V](newCap)
	for _, entry := range m.table {
		if entry != nil {
			newMap.Put(entry.key, entry.val)
		}
	}
	m.table = newMap.table
}

func main() {
	mapInstance := NewMyLinearProbingHashMap1[int, int]()
	mapInstance.Put(1, 1)
	mapInstance.Put(2, 2)
	mapInstance.Put(10, 10)
	mapInstance.Put(20, 20)
	mapInstance.Put(30, 30)
	mapInstance.Put(3, 3)

	val, _ := mapInstance.Get(1)
	fmt.Println(val) // 1
	val, _ = mapInstance.Get(2)
	fmt.Println(val) // 2
	val, _ = mapInstance.Get(20)
	fmt.Println(val) // 20

	mapInstance.Put(1, 100)
	val, _ = mapInstance.Get(1)
	fmt.Println(val) // 100

	mapInstance.Remove(20)
	fmt.Println(mapInstance.ContainsKey(20)) // false
	val, _ = mapInstance.Get(30)
	fmt.Println(val) // 30
}
