package problems

func checkOnesSegment(s string) bool {
	b := false
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == 49 {
			b = true
		}

		if b && s[i-1] == 48 {
			return false
		}
	}

	return true
}
