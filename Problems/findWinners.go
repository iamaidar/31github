package problems

import "slices"

func findWinners(matches [][]int) [][]int {
	res := make([][]int, 2)
	win := make(map[int]int)
	los := make(map[int]int)

	for _, sub := range matches {
		win[sub[0]]++
		los[sub[1]]++
	}

	for k, _ := range win {
		if _, ok := los[k]; !ok {
			res[0] = append(res[0], k)
		}
	}

	for k, v := range los {
		if v == 1 {
			res[1] = append(res[1], k)
		}
	}

	slices.Sort(res[0])
	slices.Sort(res[1])

	return res
}
