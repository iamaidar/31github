package problems

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for _, val := range nums {
		if val == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}

	return max
}
