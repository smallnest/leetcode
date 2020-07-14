package main

func main() {

}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}

	//单调栈
	mono_stack := []int{}
	// 从头开始遍历
	for i := 0; i < n; i++ {
		// 如果栈内有数据，并且栈顶的值比当前的值大, 破坏了单调,
		// 那么找到左边比当前值小的第一个值
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i    //保存比栈顶的的值小的第一个右边的位置
			mono_stack = mono_stack[:len(mono_stack)-1] // 弹出栈顶
		}

		// 记录比当前值小的第一个值的位置
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}

		// 将当前的值押入到栈中
		mono_stack = append(mono_stack, i)
	}
	ans := 0

	// 已经得到了每个元素的左边界和右边界，左右边界相减，减1乘以高度就是这个元素所在的矩形的最大面积
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
