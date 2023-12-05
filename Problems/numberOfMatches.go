package problems

func numberOfMatches(n int) int {
	var total int

	for n != 1 {
		if n%2 == 0 {
			n /= 2
			total += n
			continue
		}

		n = (n-1)/2 + 1
		total += n - 1
	}

	return total
}
