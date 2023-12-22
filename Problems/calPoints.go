package problems

func calPoints(operations []string) int {
	var stack []int

	for _, val := range operations {
		num, err := strconv.Atoi(val)
		if err != nil {
			switch val {
			case "+":
				stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			case "C":
				stack = stack[:len(stack)-1]
			case "D":
				stack = append(stack, stack[len(stack)-1]*2)
			}
			continue
		}

		stack = append(stack, num)
	}

	sum := 0

	for j := 0; j < len(stack); j++ {
		sum += stack[j]
	}

    return sum 
}
