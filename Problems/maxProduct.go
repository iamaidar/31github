package problems

func maxProduct(nums []int) int {
	var max int
	i, j := 0, len(nums)-1
	for i < j {
		temp := (nums[i] - 1) * (nums[j] - 1)
		if temp > max {
			max = temp
		}

		if nums[i] >= nums[j] {
			j--
		} else {
			i++
		}
	}

	return max
}
