package leetcode

type IntStack []int

func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *IntStack) Top() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		return (*s)[len(*s)-1], true
	}
}

func (s *IntStack) Push(e int) {
	*s = append(*s, e)
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
