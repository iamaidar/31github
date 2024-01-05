package problems

func lengthOfLIS(nums []int) int {
	temp := make([]int, len(nums))
	m := 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				temp[i] = max(temp[i], temp[j]+1)
			}
		}

		m = max(temp[i], m)
	}

	return m + 1
}
