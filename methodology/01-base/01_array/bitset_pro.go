package main

import (
	"fmt"
)

// BitMap 结构体定义
type BitMap struct {
	// bits: 底层存储，使用 uint64 数组，每个元素包含 64 个位
	bits []uint64
	// max: 能够存储的最大数字范围（用于边界检查）
	max int
}

// NewBitMap 初始化函数
// max 参数表示你需要存储的最大数字是多少
func NewBitMap(max int) *BitMap {
	// 计算需要多少个 uint64 才能存下 max 个位
	// 比如 max=64，需要 1 个 uint64；max=65，需要 2 个
	// (max + 64 - 1) / 64 是向上取整的高效写法
	size := (max + 64 - 1) / 64
	return &BitMap{
		bits: make([]uint64, size),
		max:  max,
	}
}

// Add 添加元素（将对应的位置为 1）
func (b *BitMap) Add(num int) {
	if num < 0 || num > b.max {
		return
	}

	// 1. 计算数组索引：num / 64
	// 使用位运算 >> 6 效率更高，等同于除以 64
	index := num >> 6

	// 2. 计算位偏移：num % 64
	// 使用位运算 & 63 效率更高，等同于对 64 取模
	pos := num & 63

	// 3. 将对应位置 1
	// 1 << pos 将 1 左移 pos 位
	// 比如 pos=2，二进制变成 ...00100
	// 使用 | (或) 操作，保留原有位，仅将目标位置 1
	b.bits[index] |= 1 << pos
}

// Remove 删除元素（将对应的位置为 0）
func (b *BitMap) Remove(num int) {
	if num < 0 || num > b.max {
		return
	}

	index := num >> 6
	pos := num & 63

	// 3. 将对应位置 0
	// ^(1 << pos) 是取反操作
	// 比如 1<<2 是 ...00100，取反后是 ...11011
	// 使用 & (与) 操作，目标位变成 0，其他位保持不变
	b.bits[index] &= ^(1 << pos) // Go语言中 ^ 是按位取反
}

// Has 判断元素是否存在
func (b *BitMap) Has(num int) bool {
	if num < 0 || num > b.max {
		return false
	}

	index := num >> 6
	pos := num & 63

	// 检查对应位是否为 1
	// 如果结果不等于 0，说明该位被设置过
	return (b.bits[index] & (1 << pos)) != 0
}

// String 调试输出（仅打印存在的数字）
func (b *BitMap) String() string {
	var exists []int
	// 遍历所有可能的数字，检查是否存在
	// 注意：实际生产中不会这样遍历，这里仅为演示
	for i := 0; i <= b.max; i++ {
		if b.Has(i) {
			exists = append(exists, i)
		}
	}
	return fmt.Sprintf("BitMap contains: %v", exists)
}

// --- 测试流程 ---
func main() {
	// 初始化一个能存到 100 的位图
	bm := NewBitMap(100)

	fmt.Println("1. 添加数据 5, 10, 63, 64")
	bm.Add(5)
	bm.Add(10)
	bm.Add(63) // 第0个 uint64 的最后一位
	bm.Add(64) // 第1个 uint64 的第0位

	fmt.Println(bm)

	fmt.Println("2. 查询是否存在")
	fmt.Printf("Has 10? %v\n", bm.Has(10))
	fmt.Printf("Has 20? %v\n", bm.Has(20))

	fmt.Println("3. 删除数据 10")
	bm.Remove(10)
	fmt.Printf("Has 10? %v\n", bm.Has(10))

	fmt.Println("4. 最终状态:", bm)
}
