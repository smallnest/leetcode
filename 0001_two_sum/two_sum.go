package main

func main() {

}

// 从整数数组中找到两个数，相加得到目标值.
// TAG: 哈希表, O(n)
// 思路: 对扫过的数据通过一个哈希表记录下来，方便后续的数查找
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := m[complement]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}

	return nil
}
