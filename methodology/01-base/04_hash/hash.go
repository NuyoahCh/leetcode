package main

import (
	"fmt"
	"hash/fnv"
)

// KVNode01 链表节点：存储具体的键值对
type KVNode01[V any] struct {
	Key   string
	Value V
	Next  *KVNode01[V]
}

// HashTable 哈希表结构
type HashTable[V any] struct {
	buckets []*KVNode01[V] // 桶数组，每个桶是一个链表的头节点
	size    int            // 桶的数量（物理长度）
}

// NewHashTable 初始化
// size 建议为 2 的幂次方，便于通过位运算加速（虽然这里为了逻辑简单用了取模）
func NewHashTable[V any](size int) *HashTable[V] {
	return &HashTable[V]{
		buckets: make([]*KVNode01[V], size),
		size:    size,
	}
}

// simpleHash 一个简单的哈希函数
// 将字符串转换为 0 ~ size-1 之间的整数索引
func (ht *HashTable[V]) getIndex(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	hashVal := h.Sum32()
	// 取模操作，确保落在数组范围内
	return int(hashVal) % ht.size
}

// Put 存入数据：O(1) 平均
func (ht *HashTable[V]) Put(key string, value V) {
	index := ht.getIndex(key)

	// 1. 获取该索引位置的链表头
	head := ht.buckets[index]

	// 2. 遍历链表，查看 Key 是否已存在（更新操作）
	current := head
	for current != nil {
		if current.Key == key {
			current.Value = value // 覆盖旧值
			return
		}
		current = current.Next
	}

	// 3. 如果 Key 不存在，使用“头插法”插入新节点
	// 头插法比尾插法快，因为不需要遍历到链表末尾
	newNode := &KVNode01[V]{
		Key:   key,
		Value: value,
		Next:  head, // 新节点的 Next 指向原本的头
	}

	// 4. 更新桶的头指针
	ht.buckets[index] = newNode
}

// Get 读取数据：O(1) 平均
func (ht *HashTable[V]) Get(key string) (V, bool) {
	var zero V
	index := ht.getIndex(key)

	// 遍历该桶对应的链表
	current := ht.buckets[index]
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return zero, false
}

// Remove 删除数据：O(1) 平均
func (ht *HashTable[V]) Remove(key string) {
	index := ht.getIndex(key)
	current := ht.buckets[index]

	var prev *KVNode01[V] // 记录前驱节点，用于删除

	for current != nil {
		if current.Key == key {
			// 找到目标，开始删除
			if prev == nil {
				// 删除的是头节点
				ht.buckets[index] = current.Next
			} else {
				// 删除的是中间或尾部节点
				prev.Next = current.Next
			}
			return
		}
		// 指针后移
		prev = current
		current = current.Next
	}
}

// String 调试打印
func (ht *HashTable[V]) String() string {
	result := ""
	for i, bucket := range ht.buckets {
		if bucket != nil {
			result += fmt.Sprintf("Bucket %d: ", i)
			curr := bucket
			for curr != nil {
				result += fmt.Sprintf("[%s:%v] -> ", curr.Key, curr.Value)
				curr = curr.Next
			}
			result += "nil\n"
		}
	}
	return result
}

func main() {
	// 创建一个比较小的哈希表，故意制造冲突以便观察
	ht := NewHashTable[string](5)

	fmt.Println("1. 插入数据")
	ht.Put("name", "Gemini")
	ht.Put("age", "2")
	ht.Put("lang", "Go")
	// 故意构造可能冲突的数据（取决于哈希函数，但在小桶中概率很高）
	ht.Put("key1", "data1")
	ht.Put("key2", "data2")

	fmt.Printf("%s\n", ht)

	fmt.Println("2. 查询数据")
	if val, ok := ht.Get("name"); ok {
		fmt.Println("Found name:", val)
	}

	fmt.Println("3. 更新数据")
	ht.Put("name", "Google Gemini")
	val, _ := ht.Get("name")
	fmt.Println("Updated name:", val)

	fmt.Println("4. 删除数据")
	ht.Remove("age")
	fmt.Printf("%s\n", ht)
}
