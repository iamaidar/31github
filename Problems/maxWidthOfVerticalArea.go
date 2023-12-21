package problems

import "slices"

func maxWidthOfVerticalArea(points [][]int) int {
	x := make([]int, len(points))

	max := -1

	for i := 0; i < len(points); i++ {
		x[i] = points[i][0]
	}

	slices.Sort(x)
	for i := 0; i < len(x)-1; i++ {
		if x[i+1]-x[i] > max {
			max = x[i+1] - x[i]
		}
	}
	return max
}
