package problems

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0

	for i < len([]rune(s)) && j < len([]rune(t)) {
		if []rune(s)[i] == []rune(t)[j] {
			i++
			j++
		} else {
			j++
		}
	}

	if i == len([]rune(s)) {
		return true
	}
	return false
}
