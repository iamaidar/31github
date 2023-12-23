package problems

import "strings"

//my solution
func gcdOfStrings(str1 string, str2 string) string {
	var diff string = str1

	for i := len(diff) - 1; i >= 0; i++ {
		if strings.ReplaceAll(str1, diff, "") == "" && strings.ReplaceAll(str2, diff, "") == "" {
			return diff
		} else {
			if len(diff) != 0 {
				diff = string([]rune(diff)[:len(diff)-1])
			} else {
				return ""
			}
		}
	}

	return diff
}

//one of the best solutions

// func gcdOfStrings(str1 string, str2 string) string {
//     if str1 + str2 != str2 + str1 {
//         return ""
//     }

//     a, b := len(str1), len(str2)
//     for b != 0 {
//         a, b = b, a % b
//     }

//     return str1[:a]
// }
