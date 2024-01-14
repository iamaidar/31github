package problems

import "slices"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	n, m := make(map[byte]int), make(map[byte]int)
	arr1 := []int{}
	arr2 := []int{}

	for i := 0; i < len(word1); i++ {
		m[word1[i]]++
		n[word2[i]]++
	}

	for k, v := range m {
		if _, ok := n[k]; !ok {
			return false
		}
		arr1 = append(arr1, v)
	}
	for _, v := range n {
		arr2 = append(arr2, v)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	for i := len(arr1) - 1; i >= 0; i-- {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
