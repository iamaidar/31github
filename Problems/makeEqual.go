package problems

//FROM SOLUTION(GUT GUT)
func makeEqual(words []string) bool {
	temp := make([]int, 26)

	for _, word := range words {
		for _, val := range word {
			temp[val-'a']++
		}
	}

	for _, val := range temp {
		if val%len(words) != 0 {
			return false
		}
	}

	return true
}

//MY SOLUTION
// func makeEqual(words []string) bool {
// 	m := make(map[rune]int)
// 	var temp []rune

// 	for _, word := range words {
// 		temp = append(temp, []rune(word)...)
// 	}

// 	for _, r := range temp {
// 		m[r]++
// 	}

// 	for _, val := range m {
// 		if val%len(words) != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
