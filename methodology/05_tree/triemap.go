package main

// TrieMap 底层用 Trie 树实现的键值映射
// 键为 string 类型，值为类型 V
type TrieMap struct{}

// 在 Map 中添加 key
func (tm *TrieMap) put(key string, val interface{}) {
}

// 删除键 key 以及对应的值
func (tm *TrieMap) remove(key string) {
}

// 搜索 key 对应的值，不存在则返回 nil
// get("the") -> 4
// get("tha") -> nil
func (tm *TrieMap) get(key string) interface{} {
	return nil
}

// 判断 key 是否存在在 Map 中
// containsKey("tea") -> false
// containsKey("team") -> true
func (tm *TrieMap) containsKey(key string) bool {
	return false
}

// 在 Map 的所有键中搜索 query 的最短前缀
// shortestPrefixOf("themxyz") -> "the"
func (tm *TrieMap) shortestPrefixOf(query string) string {
	return ""
}

// 在 Map 的所有键中搜索 query 的最长前缀
// longestPrefixOf("themxyz") -> "them"
func (tm *TrieMap) longestPrefixOf(query string) string {
	return ""
}

// 搜索所有前缀为 prefix 的键
// keysWithPrefix("th") -> ["that", "the", "them"]
func (tm *TrieMap) keysWithPrefix(prefix string) []string {
	return nil
}

// 判断是和否存在前缀为 prefix 的键
// hasKeyWithPrefix("tha") -> true
// hasKeyWithPrefix("apple") -> false
func (tm *TrieMap) hasKeyWithPrefix(prefix string) bool {
	return false
}

// 通配符 . 匹配任意字符，搜索所有匹配的键
// keysWithPattern("t.a.") -> ["team", "that"]
func (tm *TrieMap) keysWithPattern(pattern string) []string {
	return nil
}

// 通配符 . 匹配任意字符，判断是否存在匹配的键
// hasKeyWithPattern(".ip") -> true
// hasKeyWithPattern(".i") -> false
func (tm *TrieMap) hasKeyWithPattern(pattern string) bool {
	return false
}

// 返回 Map 中键值对的数量
func (tm *TrieMap) size() int {
	return 0
}
