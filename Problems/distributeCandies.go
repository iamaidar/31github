package problems

func distributeCandies(candyType []int) int {
	m := make(map[int]struct{})

	for _, val := range candyType {
		m[val] = struct{}{}
		if len(m) > len(candyType)/2 {
			return len(candyType) / 2
		}
	}

	if len(m) < len(candyType)/2 {
		return len(m)
	}

	return len(candyType) / 2
}
