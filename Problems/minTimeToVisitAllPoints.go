package problems

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0.0

	for i := 0; i < len(points)-1; i++ {
		x1, y1 := float64(points[i][0]), float64(points[i][1])
		x2, y2 := float64(points[i+1][0]), float64(points[i+1][1])

		if math.Abs(x2-x1) > math.Abs(y2-y1) {
			time += math.Abs(x2 - x1)
		} else {
			time += math.Abs(y2 - y1)
		}
	}

	return int(time)
}
