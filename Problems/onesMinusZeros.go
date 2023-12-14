package problems

func onesMinusZeros(grid [][]int) [][]int {
	onesRow := make([]int, len(grid))
	onesCol := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		count := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
		onesRow[i] = count
	}

	for i := 0; i < len(grid[0]); i++ {
		count := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				count++
			}
		}
		onesCol[i] = count

	}

	diff := grid

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			diff[i][j] = onesRow[i] + onesCol[j] - (len(grid) - onesRow[i]) - (len(grid[0]) - onesCol[j])
		}
	}

	return diff
}
