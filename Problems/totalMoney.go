package problems

func totalMoney(n int) int {
	var sum int
	pon := 0
	money := 0

	for i := 0; i < n; i++ {
		if i%7 == 0 {
			pon++
			money = 0
		}
		sum += pon + money
		money++
	}

	return sum
}
