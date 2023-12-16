package problems

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool)
	var res []int
	for i := 0; i < len(nums); i++ {
		m[i+1] = true
	}

	for _, val := range nums {
		m[val] = false
	}

	for key, val := range m {
		if val {
			res = append(res, key)
		}
	}

	return res

}
