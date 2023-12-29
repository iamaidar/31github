package problems

//FROM SOLUTIONS(ПОЧТИ ПОХОЖИЙ)
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prod := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prod
		prod *= nums[i]
	}

	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= prod
		prod *= nums[i]
	}

	return res
}

//MY SOLUTION
// func productExceptSelf(nums []int) []int {
// 	leftProd := make([]int, len(nums))
// 	rightProd := make([]int, len(nums))
// 	prod := 1
// 	for i := 0; i < len(nums); i++ {
// 		leftProd[i] = prod
// 		prod *= nums[i]
// 	}

// 	prod = 1

// 	for i := len(nums) - 1; i >= 0; i-- {
// 		rightProd[i] = prod
// 		prod *= nums[i]
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		nums[i] = leftProd[i] * rightProd[i]
// 	}

//     return nums
// }
