package main

func fastSlowTemplateArray(nums []int) int {
	slow := 0
	// fast 指针用于探路
	for fast := 0; fast < len(nums); fast++ {
		// 逻辑判断：如果 nums[fast] 满足保留条件（例如不等于 val，或者不重复）
		if checkCondition(nums[fast]) {
			// 将符合条件的元素放入 slow 位置
			nums[slow] = nums[fast]
			// slow 指针前进一步，扩展有效区
			slow++
		}
	}
	// 返回有效数组的长度
	return slow
}

func checkCondition(num int) bool {
	return true
}
