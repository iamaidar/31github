package problems

import "sort"

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)

	for _, val := range strs {
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		m[string(runes)] = append(m[string(runes)], val)
	}

	for _, val := range m {
		res = append(res, val)
	}

	return res
}
