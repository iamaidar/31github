package problems

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	i, j := 0, n

	for i <= j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}

	for i := 0; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
