package problems

func checkAlmostEquivalent(word1 string, word2 string) bool {
	m := make(map[rune]int)

	for _, val := range word1 {
		m[val]++
	}


	for _, val := range word2 {
		m[val]--
		if m[val] < -3 {
			return false
		}
	}

	for _, val := range m {
		if val > 3 {
			return false
		}
	}

    return true
}
