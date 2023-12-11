package problems

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	count := 0
	for _, val := range hours {
		if val >= target {
			count++
		}
	}

	return count
}
