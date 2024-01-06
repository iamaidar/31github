package problems

func tictactoe(moves [][]int) string {
	var table [][]int
	for i := 0; i < 3; i++ {
		temp := make([]int, 3)
		table = append(table, temp)
	}

	for i, v := range moves {
		if i%2 == 0 {
			table[v[0]][v[1]] = 1
		} else {
			table[v[0]][v[1]] = 2
		}

		if table[v[0]][0] == table[v[0]][1] && table[v[0]][1] == table[v[0]][2] && table[v[0]][v[1]] != 0 {
			if i%2 == 0 {
				return "A"
			} else {
				return "B"
			}
		} else if table[0][v[1]] == table[1][v[1]] && table[1][v[1]] == table[2][v[1]] && table[v[0]][v[1]] != 0 {
			if i%2 == 0 {
				return "A"
			} else {
				return "B"
			}
		} else if v[0] == v[1] && table[0][0] == table[1][1] && table[1][1] == table[2][2] && table[0][0] != 0 {
			if i%2 == 0 {
				return "A"
			} else {
				return "B"
			}
		} else if v[0]+v[1] == 2 && table[0][2] == table[1][1] && table[1][1] == table[2][0] && table[0][0] != 0 {
			if i%2 == 0 {
				return "A"
			} else {
				return "B"
			}
		}
	}

	if len(moves) != 9 {
		return "Pending"
	}

	return "Draw"
}
