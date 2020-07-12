package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var rt strings.Builder

	var min int = int(math.MaxInt64)
	var lens = make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		lens[i] = len(strs[i])
		if lens[i] < min {
			min = lens[i]
		}
	}

	for i := 0; i < min; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				return rt.String()
			}
		}
		rt.WriteByte(c)
	}

	return rt.String()
}
