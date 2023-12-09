package problems

func reverse(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}

func finalString(s string) string {
	var res string
	for _, ch := range s {
		if ch == 'i' {
			res = reverse(res)
			continue
		}
		res += string(ch)
	}
	return res
}
