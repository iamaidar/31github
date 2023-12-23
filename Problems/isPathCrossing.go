package problems

import "fmt"

func isPathCrossing(path string) bool {
	m := make(map[string]struct{})
	m["0 0"] = struct{}{}
	x, y := 0, 0
	for _, val := range path {
		switch val {
		case 'N':
			x++
		case 'E':
			y++
		case 'S':
			x--
		case 'W':
			y--
		}
		s := fmt.Sprintf("%d %d", x, y)
		if _, ok := m[s]; ok {
			return true
		}
		m[s] = struct{}{}

	}
	return false
}
