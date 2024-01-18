package problems

//Оказалось это просто числа Фиббоначи
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	a, b := 1, 2
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return a
}
