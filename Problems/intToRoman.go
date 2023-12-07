package problems

func intToRoman(num int) string {
	i := 0
	res := ""
	for ; num > 0; num /= 10 {
		n := num % 10
		i++
		switch n {
		case 1:
			if i == 1 {
				res = "I" + res
			} else if i == 2 {
				res = "X" + res
			} else if i == 3 {
				res = "C" + res
			} else {
				res = "M" + res
			}
		case 2:
			if i == 1 {
				res = "II" + res
			} else if i == 2 {
				res = "XX" + res
			} else if i == 3 {
				res = "CC" + res
			} else {
				res = "MM" + res
			}
		case 3:
			if i == 1 {
				res = "III" + res
			} else if i == 2 {
				res = "XXX" + res
			} else if i == 3 {
				res = "CCC" + res
			} else {
				res = "MMM" + res
			}
		case 4:
			if i == 1 {
				res = "IV" + res
			} else if i == 2 {
				res = "XL" + res
			} else if i == 3 {
				res = "CD" + res
			}
		case 5:
			if i == 1 {
				res = "V" + res
			} else if i == 2 {
				res = "L" + res
			} else if i == 3 {
				res = "D" + res
			}
		case 6:
			if i == 1 {
				res = "VI" + res
			} else if i == 2 {
				res = "LX" + res
			} else if i == 3 {
				res = "DC" + res
			}
		case 7:
			if i == 1 {
				res = "VII" + res
			} else if i == 2 {
				res = "LXX" + res
			} else if i == 3 {
				res = "DCC" + res
			}
		case 8:
			if i == 1 {
				res = "VIII" + res
			} else if i == 2 {
				res = "LXXX" + res
			} else if i == 3 {
				res = "DCCC" + res
			}
		case 9:
			if i == 1 {
				res = "IX" + res
			} else if i == 2 {
				res = "XC" + res
			} else if i == 3 {
				res = "CM" + res
			}
		}
	}
	return res
}
