package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	cur := 1
	start := 0

	for i := 1; i < len(s); i++ {
		c := s[i]

		// check repeated char
		found := false
		for j := i - 1; j >= start; j-- {
			if s[j] == c {
				cur = i - j
				start = j + 1
				found = true
				break
			}
		}

		// step
		if !found {
			cur++
		}
		if cur > max {
			max = cur
		}
	}

	return max
}
