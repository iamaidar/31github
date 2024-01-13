package problems

func minSteps(s string, t string) int {
	m := make([]int, 26)

	for i := 0; i < len(s); i++ {
		m[t[i]-'a']++
		m[s[i]-'a']--
	}

	count := 0

	for _, v := range m {
		if v < 0 {
			count += v
		}
	}

	return count * -1
}
