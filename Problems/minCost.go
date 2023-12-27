package problems

func minCost(colors string, neededTime []int) int {
	max := func(idx []int) (int, int) {
		m := neededTime[idx[0]]
		for _, val := range idx {
			if neededTime[val] > m {
				m = neededTime[val]
			}
		}
		count := 0
		for _, val := range idx {
			if neededTime[val] == m {
				count++
			}
		}
		return m, count
	}

	var res [][]int

	var temp []int
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			temp = append(temp, i)
		} else {
			temp = append(temp, i)
			res = append(res, temp)
			temp = []int{}
		}

		if i == len(colors)-2 && colors[i] == colors[i+1] {
			temp = append(temp, i+1)
			res = append(res, temp)
		}
	}

	time := 0

	for i := 0; i < len(res); i++ {
		if len(res[i]) > 1 {
			m, count := max(res[i])
			for j := 0; j < len(res[i]); j++ {
				if neededTime[res[i][j]] < m {
					time += neededTime[res[i][j]]
				}
			}
			time += (count - 1) * m
		}
	}

	return time
}
