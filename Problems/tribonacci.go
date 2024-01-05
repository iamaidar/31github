package problems

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	arr := make([]int, n+2)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 1

	if n > 2 {
		for i := 3; i < n+1; i++ {
			arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
		}
	} else {
		return arr[n]
	}

	return arr[n]
}
