package problems

func intersection(nums [][]int) []int {
	m := make(map[int]int)
	var res []int

	for _, val := range nums {
		for _, el := range val {
			m[el]++
			if m[el] == len(nums) {
				res = append(res, el)
			}
		}
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res); j++ {
			if res[i] < res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	return res
}
