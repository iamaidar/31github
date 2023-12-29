package problems

//GENIUS EPTA
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

//My Solution with extra space
// func rotate(nums []int, k int)  {
//     res := make([]int, len(nums))

// 	for i := range nums {
// 		res[(i+k)%len(nums)] = nums[i]
// 	}

//     copy(nums, res)
// }
