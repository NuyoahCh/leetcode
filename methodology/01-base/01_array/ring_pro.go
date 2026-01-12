package main

import "errors"

// CycleArray is a cyclic array data structure
type CycleArray[T any] struct {
	arr   []T
	start int
	end   int
	count int
	size  int
}

func NewCycleArray[T any]() *CycleArray[T] {
	return NewCycleArrayWithSize[T](1)
}

func NewCycleArrayWithSize[T any](size int) *CycleArray[T] {
	return &CycleArray[T]{
		arr:   make([]T, size),
		start: 0,
		end:   0,
		count: 0,
		size:  size,
	}
}

// 自动扩缩容辅助函数
func (ca *CycleArray[T]) resize(newSize int) {
	// 创建新的数组
	newArr := make([]T, newSize)
	// 将旧数组的元素复制到新数组中
	for i := 0; i < ca.count; i++ {
		newArr[i] = ca.arr[(ca.start+i)%ca.size]
	}
	ca.arr = newArr
	// 重置 start 和 end 指针
	ca.start = 0
	ca.end = ca.count
	ca.size = newSize
}

// 在数组头部添加元素，时间复杂度 O(1)
func (ca *CycleArray[T]) AddFirst(val T) {
	// 当数组满时，扩容为原来的两倍
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}
	// 因为 start 是闭区间，所以先左移，再赋值
	ca.start = (ca.start - 1 + ca.size) % ca.size
	ca.arr[ca.start] = val
	ca.count++
}

// 删除数组头部元素，时间复杂度 O(1)
func (ca *CycleArray[T]) RemoveFirst() error {
	if ca.isEmpty() {
		return errors.New("array is empty")
	}
	// 因为 start 是闭区间，所以先赋值，再右移
	ca.arr[ca.start] = *new(T)
	ca.start = (ca.start + 1) % ca.size
	ca.count--
	// 如果数组元素数量减少到原大小的四分之一，则减小数组大小为一半
	if ca.count > 0 && ca.count == ca.size/4 {
		ca.resize(ca.size / 2)
	}
	return nil
}

// 在数组尾部添加元素，时间复杂度 O(1)
func (ca *CycleArray[T]) AddLast(val T) {
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}
	// 因为 end 是开区间，所以是先赋值，再右移
	ca.arr[ca.end] = val
	ca.end = (ca.end + 1) % ca.size
	ca.count++
}

// 删除数组尾部元素，时间复杂度 O(1)
func (ca *CycleArray[T]) RemoveLast() error {
	if ca.isEmpty() {
		return errors.New("array is empty")
	}
	// 因为 end 是开区间，所以先左移，再赋值
	ca.end = (ca.end - 1 + ca.size) % ca.size
	ca.arr[ca.end] = *new(T)
	ca.count--
	// 缩容
	if ca.count > 0 && ca.count == ca.size/4 {
		ca.resize(ca.size / 2)
	}
	return nil
}

// 获取数组头部元素，时间复杂度 O(1)
func (ca *CycleArray[T]) GetFirst() (T, error) {
	if ca.isEmpty() {
		return *new(T), errors.New("array is empty")
	}
	return ca.arr[ca.start], nil
}

// 获取数组尾部元素，时间复杂度 O(1)
func (ca *CycleArray[T]) GetLast() (T, error) {
	if ca.isEmpty() {
		return *new(T), errors.New("array is empty")
	}
	// end 是开区间，指向的是下一个元素的位置，所以要减 1
	return ca.arr[(ca.end-1+ca.size)%ca.size], nil
}

func (ca *CycleArray[T]) isFull() bool {
	return ca.count == ca.size
}

func (ca *CycleArray[T]) Size() int {
	return ca.count
}

func (ca *CycleArray[T]) isEmpty() bool {
	return ca.count == 0
}
