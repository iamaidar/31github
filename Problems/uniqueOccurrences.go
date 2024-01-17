package problems

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	n := make(map[int]struct{})

	for _, val := range arr {
		m[val]++
	}

	for _, v := range m {
		if _, ok := n[v]; ok {
			return false
		} else {
			n[v] = struct{}{}
		}
	}

	return true
}
