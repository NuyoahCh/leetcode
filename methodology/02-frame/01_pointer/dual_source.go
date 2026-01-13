package main

func dualSourceTemplate(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 虚拟头结点技巧：简化新链表构建的边界处理
	dummy := &ListNode{Val: -1}
	p := dummy // 结果链表的构建指针

	p1, p2 := l1, l2 // 双头指针初始化

	// 2. 双指针遍历逻辑
	for p1 != nil && p2 != nil {
		// 比较逻辑 (例如合并有序链表)
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		// 构建指针前进
		p = p.Next
	}

	// 3. 处理剩余节点 (或连接分解后的两部分)
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}
