package problems

func numSpecial(nums [][]int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			if vert(j, nums) && gor(i, nums) && nums[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func gor(a int, nums [][]int) bool {
	count := 0
	for i := 0; i < len(nums[a]); i++ {
		if nums[a][i] == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}

	return true
}

func vert(a int, nums [][]int) bool {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i][a] == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}

	return true
}
