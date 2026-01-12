package main

import (
	"errors"
	"fmt"
)

// DynamicArray 动态数组实现
type DynamicArray[T any] struct {
	data []T
	size int
	cap  int
}

// NewDynamicArray 初始化
// 建议：如果预知数据量，initialCapacity 尽量给大，避免早期的频繁扩容
func NewDynamicArray[T any](initialCapacity int) *DynamicArray[T] {
	if initialCapacity == 0 {
		initialCapacity = 10 // 给个默认值，防止 0 容量导致无法扩容
	}
	return &DynamicArray[T]{
		data: make([]T, initialCapacity),
		size: 0,
		cap:  initialCapacity,
	}
}

// Add 添加元素
// 重点：手动实现 resize 逻辑
func (d *DynamicArray[T]) Add(element T) {
	if d.size == d.cap {
		d.resize() // 触发扩容
	}
	d.data[d.size] = element
	d.size++
}

// resize 扩容的核心逻辑：以空间换时间
func (d *DynamicArray[T]) resize() {
	// 策略：这里简单采用 2 倍扩容
	// 注：Go 官方 slice 的策略更复杂（小于阈值 2 倍，大于阈值 1.25 倍平滑过渡）
	newCap := d.cap * 2

	// 1. 申请更大的新内存块
	newData := make([]T, newCap)

	// 2. 数据搬迁 (Deep Copy)
	// Go 的 copy 是内存级别的拷贝，速度很快
	copy(newData, d.data)

	// 3. 替换引用，旧数组 data 将被 GC 回收
	d.data = newData
	d.cap = newCap

	fmt.Printf("Debug: Resized to capacity %d\n", newCap)
}

// RemoveLast 删除末尾元素
// 重点：内存泄漏的防御
func (d *DynamicArray[T]) RemoveLast() (T, error) {
	var zero T
	if d.size == 0 {
		return zero, errors.New("array is empty")
	}

	// 取出元素
	item := d.data[d.size-1]

	// 关键步骤：显式置为零值
	// 如果 T 是指针类型，不置空会导致旧对象无法被 GC，造成内存泄漏
	d.data[d.size-1] = zero

	d.size--

	// 缩容逻辑（可选）：通常采用惰性缩容
	// 比如 size < cap / 4 时，才缩容到 cap / 2，防止由于频繁增删导致的“抖动”

	return item, nil
}

func (d *DynamicArray[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= d.size {
		return zero, errors.New("index out of bounds")
	}
	return d.data[index], nil
}

// 主函数测试
func main() {
	fmt.Println("=== Static Array Demo ===")
	sa := NewStaticArray[int](2)
	sa.Add(10)
	sa.Add(20)
	err := sa.Add(30) // 这里会报错
	if err != nil {
		fmt.Println("Expected Error:", err)
	}

	fmt.Println("\n=== Dynamic Array Demo ===")
	da := NewDynamicArray[int](2)
	da.Add(100)
	da.Add(200)
	// 触发扩容：容量从 2 -> 4
	da.Add(300)

	val, _ := da.Get(2)
	fmt.Printf("Get index 2: %d (Cap: %d)\n", val, da.cap)

	da.RemoveLast()
	fmt.Printf("After RemoveLast size: %d\n", da.size)
}
