package problems

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)

	m := make(map[int]struct{})
	n := make(map[int]struct{}{})

	for _, val := range nums1 {
		m[val] = struct{}{}
	}

	for _, val := range nums2 {
		n[val] = struct{}{}
	}

	for k, _ := range m {
		if _, ok := n[k]; !ok {
			res[0] = append(res[0], k)
		}
	}

	for k, _ := range n {
		if _, ok := m[k]; !ok {
			res[1] = append(res[1], k)
		}
	}

	return res
}
