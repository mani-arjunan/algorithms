package main

func isValidSudoku(board [][]byte) bool {
	hashMap := make(map[byte]int)

	for i, _ := range board {
		for j, _ := range board[i] {
			key := board[i][j]
			_, ok := hashMap[key]

			if !ok {
				if key != '.' {
					hashMap[key] = 1
				}
			} else {
				return false
			}
		}

		hashMap = map[byte]int{}
	}

	for i, _ := range board {
		for j, _ := range board[i] {
			key := board[j][i]
			_, ok := hashMap[key]

			if !ok {
				if key != '.' {
					hashMap[key] = 1
				}
			} else {
				return false
			}
		}

		hashMap = map[byte]int{}
	}

	for k := 0; k < 9; k++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				row := (k/3)*3 + i
				col := (k%3)*3 + j
				key := board[row][col]
				_, ok := hashMap[key]
				if !ok {
					if key != '.' {
						hashMap[key] = 1
					}
				} else {
					return false
				}
			}
		}
		hashMap = map[byte]int{}
	}

	return true
}
