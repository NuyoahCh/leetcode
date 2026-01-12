package main

import (
	"errors"
	"fmt"
)

// StaticArray 静态数组实现
// 核心约束：一旦初始化，Capacity 不可变，不支持动态扩容
type StaticArray[T any] struct {
	data []T // 底层虽然用了 slice，但在逻辑上我们将它视为定长内存块
	size int // 当前实际存储的元素个数
	cap  int // 绝对容量上限
}

// NewStaticArray 初始化：一次性申请内存
func NewStaticArray[T any](capacity int) *StaticArray[T] {
	return &StaticArray[T]{
		// make 申请了固定大小的内存，类似于 C 的 malloc
		data: make([]T, capacity),
		size: 0,
		cap:  capacity,
	}
}

// Set 修改元素：O(1)
func (s *StaticArray[T]) Set(index int, element T) error {
	if index < 0 || index >= s.size {
		return errors.New("index out of bounds")
	}
	s.data[index] = element
	return nil
}

// Get 获取元素：O(1)
func (s *StaticArray[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= s.size {
		return zero, errors.New("index out of bounds")
	}
	return s.data[index], nil
}

// Add 添加元素：O(1)
// 关键点：达到容量上限直接报错，绝不扩容
func (s *StaticArray[T]) Add(element T) error {
	if s.size >= s.cap {
		return errors.New("array is full: static array cannot resize")
	}
	s.data[s.size] = element
	s.size++
	return nil
}

// 打印当前状态
func (s *StaticArray[T]) String() string {
	return fmt.Sprintf("Size: %d, Cap: %d, Data: %v", s.size, s.cap, s.data[:s.size])
}
