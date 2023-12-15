package problems

func destCity(paths [][]string) string {
	temp := paths[0][1]
	for k := 0; k < len(paths); k++ {
		for i := 0; i < len(paths); i++ {
			for j := 0; j < len(paths); j++ {
				if paths[i][0] == temp {
					temp = paths[i][1]
				}
			}
		}
	}
	return temp
}

// out := make(map[string]int)

// for _, path := range paths {
// 	out[path[0]]++
// }

// fmt.Println(out)

// for _, path := range paths {
// 	if _, ok := out[path[1]]; !ok {
// 		return path[1]
// 	}
// }
