package main

func firstPalindrome(words []string) string {
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if word[i] != word[len(word)-i-1] {
				break
			}
			if i == len(word)/2 {
				return word
			}
		}
	}

	return ""
}
