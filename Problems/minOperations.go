package problems

func minOperations(s string) int {
	count := 0
	min := 0
	r := []rune(s)

	for i := 0; i < len(r)-1; i++ {
		if r[i] == '0' && r[i+1] != '1' {
			r[i+1] = '1'
			count++
		} else if r[i] == '1' && r[i+1] != '0' {
			r[i+1] = '0'
			count++
		}
	}

	r = []rune(s)

	if r[0] == '0' {
		r[0] = '1'
		min++
	} else {
		r[0] = '0'
		min++
	}

	for i := 0; i < len(r)-1; i++ {
		if r[i] == '0' && r[i+1] != '1' {
			r[i+1] = '1'
			min++
		} else if r[i] == '1' && r[i+1] != '0' {
			r[i+1] = '0'
			min++
		}
	}

	if count < min {
		return count
	}

	return min
}
