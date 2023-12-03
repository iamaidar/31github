package problems

import "math"

func countKDifference(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) == float64(k) {
				count++
			}
		}
	}
	return count
}
