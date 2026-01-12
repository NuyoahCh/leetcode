package main

import (
	"errors"
	"fmt"
)

// Stack 基于切片实现的栈
type Stack[T any] struct {
	items []T
}

// Push 入栈：O(1)
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 出栈：O(1)
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}

	// 1. 获取栈顶元素（切片最后一个元素）
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]

	// 2. 关键工程细节：内存清理
	// 如果 T 是指针或包含指针的结构体，必须显式置空
	// 否则底层的数组会一直引用该对象，导致 GC 无法回收（内存泄漏）
	s.items[lastIndex] = zero

	// 3. 调整切片长度
	s.items = s.items[:lastIndex]

	return item, nil
}

// Peek 查看栈顶元素（不删除）：O(1)
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty 判空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack[string]{}

	stack.Push("Google")
	stack.Push("OpenAI")
	stack.Push("DeepMind") // 栈顶

	val, _ := stack.Pop()
	fmt.Println("Popped:", val) // DeepMind

	top, _ := stack.Peek()
	fmt.Println("Current Top:", top) // OpenAI
}
