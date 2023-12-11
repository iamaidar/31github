package problems

func findSpecialInteger(arr []int) int {
	m := make(map[int]int)

	for _, val := range arr {
		m[val]++
		if m[val] > len(arr)/4 {
			return val
		}
	}

	return 0
}
