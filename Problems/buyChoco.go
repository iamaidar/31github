package problems

func buyChoco(prices []int, money int) int {
	s1, s2 := 99999999, 999999999
	for i := 0; i < len(prices); i++ {
		if prices[i] < s1 {
			s2 = s1
			s1 = prices[i]
		} else if prices[i] < s2 {
			s2 = prices[i]
		}
	}
	if s1+s2 > money {
		return money
	}
	return money - s1 - s2
}
