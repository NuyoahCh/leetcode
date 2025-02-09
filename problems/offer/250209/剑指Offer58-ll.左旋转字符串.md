# 官方链接

题目已经修改

## 题解

题目：给你一个字符串和一个k，把字符串前k个字符截取放到字符串未尾。abcd 2.-----cdab
思路：直接操作即可，easy

```go
func reverseLeftWords(s string, n int) string {
  return s[n:]+s[:n]
}
```



