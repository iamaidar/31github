package problems

func lemonadeChange(bills []int) bool {
	n := make([]int, 2)

	for _, val := range bills {
		if val == 5 {
			n[0]++
			continue
		}
		if val == 10 {
			n[1]++
			if n[0] <= 0 {
				return false
			} else {
				n[0]--
			}
		} else if val == 20 {
			if n[1] <= 0 && n[0] <= 0 {
				return false
			} else if n[1] > 0 && n[0] > 0 {
				n[1]--
				n[0]--
				continue
			}

			if n[0] < 3 {
				return false
			} else {
				n[0] -= 3
			}
		}
	}

	return true
}
