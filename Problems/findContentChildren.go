package problems

import "slices"

func findContentChildren(g []int, s []int) int {
	count := 0
	slices.Sort(g)
	slices.Sort(s)

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] >= g[i] {
				count++
				s[j] = -1
				break
			}
		}
	}

	return count
}
