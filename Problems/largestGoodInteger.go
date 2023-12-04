package problems

func largestGoodInteger(num string) string {
	res := ""
	var max int
	var id int = -1
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if int(num[i]+num[i+1]+num[i+1]) >= max {
				max = int(num[i] + num[i+1] + num[i+1])
				id = i
			}
		}
	}
	if id != -1 {
		res = string(num[id]) + string(num[id+1]) + string(num[id+2])
	}
	return res
}
