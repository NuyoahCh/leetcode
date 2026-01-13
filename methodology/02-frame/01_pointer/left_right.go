package main

func leftRightTemplate(nums []int, target int) {
	left, right := 0, len(nums)-1

	for left < right {
		// 逻辑 A：根据题目要求计算当前值（如 sum = nums[left] + nums[right]）
		current := calculate(nums[left], nums[right])

		if current == target {
			// 找到目标
			return
		} else if current < target {
			// 比如两数之和偏小，因为数组有序，需要让左指针右移变大
			left++
		} else {
			// 比如两数之和偏大，需要让右指针左移变小
			right--
		}
	}
}

func calculate(num1, num2 int) int {
	return num1 + num2
}
