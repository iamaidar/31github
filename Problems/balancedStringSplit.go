package problems

func balancedStringSplit(s string) int {
	res := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}

		if count == 0 {
			res++
		}
	}

	return res
}
