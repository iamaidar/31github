package problems

import "slices"

func removeDuplicates(nums []int) int {
	var res int = len(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i+2] == nums[i+1] {
			nums[i] = 9999999
			res--
		}
	}

	slices.Sort(nums)
	return res
}
