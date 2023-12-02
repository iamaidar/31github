package problems

import "strconv"

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		bin := strconv.FormatInt(int64(i), 2)
		count := 0
		for _, ch := range bin {
			if ch == []rune("1")[0] {
				count++
			}
		}
		if count == k {
			sum += nums[i]
		}
	}
	return sum
}
