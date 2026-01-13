package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func fastSlowTemplate(head *ListNode) *ListNode {
	// 初始化快慢指针
	slow, fast := head, head

	// 1. 针对“倒数第 k 个节点”的特殊初始化
	// for i := 0; i < k; i++ { fast = fast.Next }

	// 遍历链表
	for fast != nil && fast.Next != nil {
		// 逻辑 A：寻找中点/判断环
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步

		// 逻辑 B：判断环是否存在
		if slow == fast {
			// 进阶：寻找环起点逻辑
			// slow = head
			// while slow != fast { slow++, fast++ }
			return slow
		}
	}

	// 逻辑 C：针对“倒数第 k 个节点”的后续遍历
	// while fast != nil { slow++, fast++ }

	return slow // 返回目标节点
}
