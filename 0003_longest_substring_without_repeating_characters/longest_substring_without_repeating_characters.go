package main

func main() {

}

// 不带重复字母的最长子串. (子字符串substring，不是子序列subsequence)
// TAG: O(n)
// 思路: 使用数组记录已发现字符的当前位置。使用两个指针start, i移动，当前的搜寻范围为(start,i)的字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var m = make(map[byte]int, 256)

	maxLen := 0
	start := -1
	for i := 0; i < len(s); i++ {
		if j, ok := m[s[i]]; ok {
			if j > start { // 在本次的搜寻范围内
				start = j
			}
		}
		m[s[i]] = i
		if i-start > maxLen {
			maxLen = i - start
		}
	}
	return maxLen
}
