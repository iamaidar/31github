package problems

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0
	for _, val := range []rune(fmt.Sprintf("%b", num)) {
		if val == []rune("1")[0] {
			count++
		}
	}
	return count
}
