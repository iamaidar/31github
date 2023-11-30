package problems

func findWordsContaining(words []string, x byte) []int {
	var res []int

	for idx, word := range words {
		var check bool
		for _, val := range []rune(word) {
			if []rune(string(x))[0] == val {
				check = true
			}
		}
		if check {
			res = append(res, idx)
		}
	}
	return res
}
