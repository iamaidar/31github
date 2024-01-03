package problems

func numberOfBeams(bank []string) int {
	count := 0
	temp := 0
	temp2 := 0

	for i := 0; i < len(bank)-1; i++ {
		for j := 0; j < len(bank[i+1]); j++ {
			if bank[i][j] == '1' {
				temp++
			}
			if bank[i+1][j] == '1' {
				temp2++
			}
		}

		count += temp * temp2
		if temp2 != 0 {
			temp = 0
			temp2 = 0
		}
	}

	return count
}
