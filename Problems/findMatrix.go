package problems

func findMatrix(nums []int) [][]int {
	var res [][]int
	m := make(map[int]int)
	max := 0

	for _, val := range nums {
		m[val]++
		if m[val] > max {
			max = m[val]
		}
	}

	for i := 0; i < max; i++ {
		var temp []int
		for k, val := range m {
			if val != 0 {
				m[k]--
				temp = append(temp, k)
			}
		}

		res = append(res, temp)
	}

	return res
}
