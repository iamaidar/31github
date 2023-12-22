package problems

func maxScore(s string) int {
	max := 0

	for i := 1; i < len(s); i++ {
		ones := 0
		zeros := 0
		for j := 0; j < len(s); j++ {
			if s[j] == '0' && j < i {
				zeros++
			}
			if s[j] == '1' && j >= i {
				ones++
			}
		}

		if zeros+ones > max {
			max = zeros + ones
		}
	}

	return max
}
