package main

import (
	"errors"
	"fmt"
)

// RingBuffer 定义环形数组结构
type RingBuffer[T any] struct {
	data     []T // 物理存储载体
	head     int // 读指针：指向最早放入的元素
	tail     int // 写指针：指向下一个可以写入的位置
	count    int // 有效元素个数
	capacity int // 物理容量
}

// NewRingBuffer 初始化函数
// 申请固定大小的内存，生命周期内不再扩容
func NewRingBuffer[T any](cap int) *RingBuffer[T] {
	if cap <= 0 {
		cap = 10 // 默认保护
	}
	return &RingBuffer[T]{
		data:     make([]T, cap),
		head:     0,
		tail:     0,
		count:    0,
		capacity: cap,
	}
}

// Enqueue 入队操作（写入）
// 逻辑：写入数据 -> 移动尾指针 -> 增加计数
func (r *RingBuffer[T]) Enqueue(element T) error {
	if r.count == r.capacity {
		return errors.New("buffer is full")
	}

	// 1. 在尾指针位置写入数据
	r.data[r.tail] = element

	// 2. 移动尾指针（关键点：取模运算实现回绕）
	// 如果 tail 到达末尾，(tail + 1) % cap 会自动变为 0
	r.tail = (r.tail + 1) % r.capacity

	// 3. 更新计数
	r.count++
	return nil
}

// Dequeue 出队操作（读取）
// 逻辑：读取数据 -> 移动头指针 -> 减少计数
func (r *RingBuffer[T]) Dequeue() (T, error) {
	var zero T
	if r.count == 0 {
		return zero, errors.New("buffer is empty")
	}

	// 1. 获取头指针处的数据
	element := r.data[r.head]

	// 2. 为了避免内存泄漏（针对指针类型），将原位置置空
	r.data[r.head] = zero

	// 3. 移动头指针（同样使用取模运算）
	r.head = (r.head + 1) % r.capacity

	// 4. 更新计数
	r.count--
	return element, nil
}

// IsFull 判满
func (r *RingBuffer[T]) IsFull() bool {
	return r.count == r.capacity
}

// IsEmpty 判空
func (r *RingBuffer[T]) IsEmpty() bool {
	return r.count == 0
}

// String 调试输出
// 打印物理数组和逻辑指针位置，方便理解
func (r *RingBuffer[T]) String() string {
	return fmt.Sprintf("Data: %v | Head: %d | Tail: %d | Count: %d/%d",
		r.data, r.head, r.tail, r.count, r.capacity)
}

// --- 测试用例 ---
func main() {
	// 创建一个容量为 3 的环形数组
	rb := NewRingBuffer[int](3)

	fmt.Println("1. 初始状态:", rb)

	// 填满数组
	rb.Enqueue(10)
	rb.Enqueue(20)
	rb.Enqueue(30)
	fmt.Println("2. 填满后:", rb)

	// 尝试再插入，应该报错
	err := rb.Enqueue(40)
	if err != nil {
		fmt.Println("3. 溢出测试:", err)
	}

	// 消费两个元素
	val1, _ := rb.Dequeue()
	val2, _ := rb.Dequeue()
	fmt.Printf("4. 消费元素: %d, %d. 当前状态: %s\n", val1, val2, rb)

	// 关键时刻：再次插入，验证由于 Head 前移，Tail 是否能回到数组头部（回绕）
	rb.Enqueue(50)
	rb.Enqueue(60)
	fmt.Println("5. 回绕写入后:", rb)
	// 注意观察输出中 Data 的变化，你会发现索引 0 和 1 的位置被复用了
}
