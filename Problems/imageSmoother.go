package problems

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := range img {
		res[i] = make([]int, len(img[i]))
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			sum := 0
			count := 0
			for n := i - 1; n < i+2; n++ {
				for m := j - 1; m < j+2; m++ {
					if n >= 0 && m >= 0 && n < len(img) && m < len(img[0]) {
						count++
						sum += img[n][m]
					}
				}
			}
			res[i][j] = sum / count
		}
	}

	return res
}
