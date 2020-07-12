package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var l = len(s)
	var begin int
	var maxLen int

	for i := 0; i < l-maxLen/2; i++ {
		var j = i
		var k = i

		for k < l-1 && s[k] == s[k+1] {
			k++
		}

		// 注意两个端点
		for j > 0 && k < l-1 && s[j-1] == s[k+1] {
			j--
			k++
		}

		newLen := k - j + 1
		if newLen > maxLen {
			begin = j
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}
