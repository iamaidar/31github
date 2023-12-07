package problems

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i]%2 != 0 {
			return string(num[:i+1])
		}
	}

	return ""
}
