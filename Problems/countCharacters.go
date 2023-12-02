package problems

func countCharacters(words []string, chars string) int {
	freq := make(map[rune]int)
	sum := 0

	for _, ch := range chars {
		freq[ch]++
	}

	for _, word := range words {
		pal := make(map[rune]int)

		for _, ch := range word {
			pal[ch]++
		}

		f := false

		for key, _ := range pal {
			if freq[key] >= pal[key] {
				f = true
			} else {
				f = false
				break
			}
		}

		if f {
			sum += len([]rune(word))
		}
	}
	return sum
}
