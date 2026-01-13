package main

func palindromeExpand(s string, l, r int) string {
	// 防止索引越界，同时判断是否满足回文条件
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开
		l--
		r++
	}
	// 返回找到的回文子串
	return s[l+1 : r]
}
