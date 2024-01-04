package problems

func check(k int) int {
	count := 0

	for i := 0; i < k; i += 2 {
		if (k-i)%3 == 0 {
			count = (k - i) / 3
			break
		}
	}

	return count
}

func minOperations(nums []int) int {
	m := make(map[int]int)

	count := 0

	for _, val := range nums {
		m[val]++
	}

	for _, val := range m {
		if val == 1 {
			return -1
		}

		if val >= 3 {
			t := check(val)
			count += t + (val-t*3)/2
		} else {
			count += val / 2
		}
	}

	return count
}
