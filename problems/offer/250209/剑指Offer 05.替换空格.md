# 官方链接

题目已经修改

## 题解

题目：把字符串的空格替换为%20

思路：遍历，遇到空格替换即可

```go
func replaceSpace (s string) string {
  str := ""
  for _, v := range s {
    if string(v) != "" {
      str += string(s)
    }else{
      str += "%20"
    }
  }
  return str
}
```

