package problems

import "slices"

func findErrorNums(nums []int) []int {
	slices.Sort(nums)
	v := 1
	m := make([]int, 10002)

	for _, val := range nums {
		m[val]++
		if ok := m[v]; ok != 0 {
			v++
		}

		if m[val] == 2 {
			m[10001] = val
		}
	}

	return []int{m[10001], v}
}
