package main

func main() {

}

// 寻找两个已排序数组的中位数.
// TAG: O(n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)
	num := make([]int, m+n, m+n)
	t := (m + n) / 2

	var i, j, k int
	for i < m && j < n && k <= t {
		if nums1[i] <= nums2[j] {
			num[k] = nums1[i]
			i = i + 1
		} else {
			num[k] = nums2[j]
			j = j + 1
		}
		k = k + 1
	}

	for i < m && k <= t {
		num[k] = nums1[i]
		i = i + 1
		k = k + 1
	}

	for j < n && k <= t {
		num[k] = nums2[j]
		j = j + 1
		k = k + 1
	}

	if t*2 < (m + n) {
		return float64(num[t])
	}
	return float64(num[t-1]+num[t]) / 2

}
