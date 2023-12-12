package problems

import "strings"

func lengthOfLastWord(s string) int {
	str := strings.TrimRight(s, " ")
	arr1 := strings.Split(str, " ")
	sr := []rune(arr1[len(arr1)-1])
	return len(sr)
}
