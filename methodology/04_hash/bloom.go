package main

import (
	"fmt"
)

type SimpleBloomFilter struct {
	// 位图的大小
	bitSetSize int
	// 位图
	bitSet *MyBitSet
	// 哈希函数的个数
	k int
}

// NewSimpleBloomFilter 构造一个布隆过滤器，指定位图大小和哈希函数个数
func NewSimpleBloomFilter(bitSetSize, hashFunctionNum int) *SimpleBloomFilter {
	return &SimpleBloomFilter{
		bitSetSize: bitSetSize,
		k:          hashFunctionNum,
		bitSet:     NewMyBitSet(bitSetSize),
	}
}

// Add 添加元素
func (bf *SimpleBloomFilter) Add(element string) {
	// 获取 k 个不同的哈希值
	// 将这 k 个哈希值对应的位图中的位都设置为 1
	for i := 0; i < bf.k; i++ {
		hashValue := bf.hash(element, i)
		bf.bitSet.Set(hashValue)
	}
}

// Contains 判断元素是否存在
func (bf *SimpleBloomFilter) Contains(element string) bool {
	// 获取 k 个不同的哈希值
	// 检查这 k 个哈希值对应的位图中的位是否全部为 1
	for i := 0; i < bf.k; i++ {
		hashValue := bf.hash(element, i)
		if !bf.bitSet.Get(hashValue) {
			return false
		}
	}
	return true
}

// 模拟多个哈希函数，实际生产环境中应该使用更复杂的哈希算法
func (bf *SimpleBloomFilter) hash(element string, seed int) int {
	// 这里简化处理，实现一个简单的字符串哈希函数，和递增的索引作为种子来模拟多个哈希函数
	// 在实际应用中，为了减少哈希冲突，应该使用更复杂的哈希函数
	// 同时，种子也应该选择无规律的大质数，而不是简单的递增索引
	h := 0
	for _, ch := range element {
		h = 31*h + int(ch)
	}
	result := (h + seed) % bf.bitSetSize
	if result < 0 {
		result = -result
	}
	return result
}

func main() {
	// 创建一个位数组大小为 1000000，使用 3 个哈希函数的布隆过滤器
	bloomFilter := NewSimpleBloomFilter(1000000, 3)

	// 添加元素
	bloomFilter.Add("apple")
	bloomFilter.Add("banana")
	bloomFilter.Add("orange")

	// 检查元素是否存在
	fmt.Println("Contains apple:", bloomFilter.Contains("apple"))   // true
	fmt.Println("Contains banana:", bloomFilter.Contains("banana")) // true
	fmt.Println("Contains grape:", bloomFilter.Contains("grape"))   // false
}
