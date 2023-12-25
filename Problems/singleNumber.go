package problems

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
