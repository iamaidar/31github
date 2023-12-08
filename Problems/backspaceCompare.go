package problems

type Stack []byte

func (s *Stack) Push(elem byte) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() {
	if len(*s) != 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Stack) Equal(st Stack) bool {
	if len(*s) != len(st) {
		return false
	}

	for i := 0; i < len(st); i++ {
		if (*s)[i] != st[i] {
			return false
		}
	}

	return true
}
func backspaceCompare(s string, t string) bool {
	stack := Stack{}
	stack1 := Stack{}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "#" {
			stack.Pop()
			continue
		}

		stack.Push(s[i])
	}

	for i := 0; i < len(t); i++ {
		if string(t[i]) == "#" {
			stack1.Pop()
			continue
		}

		stack1.Push(t[i])
	}

	return stack.Equal(stack1)
}
