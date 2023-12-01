package problems

func differenceOfSums(n int, m int) int {
	var n1, n2 int
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			n1 += i
			continue
		}
		n2 += i
	}

	return n1 - n2
}
