package main

import "fmt"

type MyBitSet struct {
	// 使用 uint64 数组作为位图的底层存储
	words []uint64
	// 位图能够存储的最大元素值 + 1
	size int
}

func NewMyBitSet(size int) *MyBitSet {
	// 根据 size 计算需要多少个 uint64 来存储
	arraySize := size/64 + 1
	return &MyBitSet{
		words: make([]uint64, arraySize),
		size:  size,
	}
}

// Get 判断指定比特位是否为 1
func (bs *MyBitSet) Get(bitIndex int) bool {
	if bitIndex < 0 || bitIndex >= bs.size {
		panic(fmt.Sprintf("bitIndex must be between 0 and %d", bs.size-1))
	}
	// 找到 bitIndex 在 words 数组中的索引
	wordIndex := bitIndex / 64
	// 找到 bitIndex 在 uint64 值中的具体 bit 位
	bitOffset := bitIndex % 64
	// 使用 & 操作判断该位是否为 1
	return (bs.words[wordIndex] & (1 << bitOffset)) != 0
}

// Set 将指定比特位设置为 1
func (bs *MyBitSet) Set(bitIndex int) {
	if bitIndex < 0 || bitIndex >= bs.size {
		panic(fmt.Sprintf("bitIndex must be between 0 and %d", bs.size-1))
	}
	// 找到 bitIndex 在 words 数组中的索引
	wordIndex := bitIndex / 64
	// 找到 bitIndex 在 uint64 值中的具体 bit 位
	bitOffset := bitIndex % 64
	// 使用 | 操作将该位置 1
	bs.words[wordIndex] |= 1 << bitOffset
}

// Clear 将指定比特位设置为 0
func (bs *MyBitSet) Clear(bitIndex int) {
	if bitIndex < 0 || bitIndex >= bs.size {
		panic(fmt.Sprintf("bitIndex must be between 0 and %d", bs.size-1))
	}
	wordIndex := bitIndex / 64
	bitOffset := bitIndex % 64
	// 使用 & 和 ~ 操作将该位置 0
	bs.words[wordIndex] &= ^(1 << bitOffset)
}

func main() {
	bitSet := NewMyBitSet(1000)

	// cnt := []int{26}
	// slice_demo := make([]int, 0, 5)

	bitSet.Set(10)
	bitSet.Set(100)
	bitSet.Set(500)

	fmt.Println("Get 10:", bitSet.Get(10))   // true
	fmt.Println("Get 100:", bitSet.Get(100)) // true
	fmt.Println("Get 200:", bitSet.Get(200)) // false

	bitSet.Clear(100)
	fmt.Println("Get 100 after clear:", bitSet.Get(100)) // false
}
