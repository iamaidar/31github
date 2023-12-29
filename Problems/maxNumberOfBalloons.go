package problems

import "math"

//balloon

func maxNumberOfBalloons(text string) int {
	b := 0.0
	a := 0.0
	l := 0.0
	o := 0.0
	n := 0.0

	for _, val := range text {
		switch val {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l += 0.5
		case 'o':
			o += 0.5
		case 'n':
			n++
		}
	}

	if b < 1 || a < 1 || l < 1 || o < 1 || n < 1 {
		return 0
	}

	return int(math.Min(b, math.Min(a, math.Min(l, math.Min(o, n)))))
}

//NON EFFECTIVE (MY SOLUTION)
// func maxNumberOfBalloons(text string) int {
// 	m := make(map[rune]int)

// 	b := "balloon"

// 	for _, val := range text {
// 		m[val]++
// 	}

// 	min := 9999999999999

// 	for _, val := range b {
// 		if _, ok := m[val]; !ok {
// 			return 0
// 		}

// 		if (val == 'l' && m[val] < 2) || (val == 'o' && m[val] < 2) {
// 			return 0
// 		}

// 		if val == 'l' || val == 'o' {
// 			if m[val]/2 < min {
// 				min = m[val] / 2
// 			}
// 		}

// 		if m[val] < min && val != 'l' && val != '0' {
// 			min = m[val]
// 		}
// 	}

//     return min
// }
