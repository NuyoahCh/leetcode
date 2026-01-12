package main

type MyHashSet[K comparable] struct {
	// 底层 map，用于存储集合元素，值用空结构体作为占位符
	data map[K]struct{}
}

func NewMyHashSet[K comparable]() *MyHashSet[K] {
	return &MyHashSet[K]{
		data: make(map[K]struct{}),
	}
}

func (s *MyHashSet[K]) Add(key K) {
	// 向哈希表添加一个键值对，值用空结构体作为占位符
	s.data[key] = struct{}{}
}

func (s *MyHashSet[K]) Remove(key K) {
	// 从哈希表中移除键 key
	delete(s.data, key)
}

func (s *MyHashSet[K]) Contains(key K) bool {
	// 判断哈希表中是否包含键 key
	_, exists := s.data[key]
	return exists
}

func (s *MyHashSet[K]) Size() int {
	return len(s.data)
}
