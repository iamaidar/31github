package problems

//MY SOLUTION (NO OPTIMIZE)
// func maxLengthBetweenEqualCharacters(s string) int {
// 	max := -1

// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			if s[i] == s[j] {
// 				if j-i > max {
// 					max = j-i
// 				}
// 			}
// 		}
// 	}

//     max--
//     if max < -1 {
//         return -1
//     }

//     return max
// }
