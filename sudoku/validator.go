package sudoku

func checkLine(board [][]byte) bool {
	for _, line := range board {
		counter := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, item := range line {
			if item == '.' {
				continue
			}

			pos := item - 49
			if counter[pos] > 0 {
				return false
			}
			counter[pos] = 1
		}
	}
	return true
}

func checkColumn(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		counter := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for j := 0; j < 9; j++ {
			item := board[j][i]

			if item == '.' {
				continue
			}

			pos := item - 49
			if counter[pos] > 0 {
				return false
			}
			counter[pos] = 1
		}
	}
	return true
}

func checkBlock(board [][]byte) bool {
	positionMap := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			startX := i * 3
			startY := j * 3

			counter := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			for _, posMap := range positionMap {
				offsetX := posMap[0]
				offsetY := posMap[1]
				item := board[startX+offsetX][startY+offsetY]

				if item == '.' {
					continue
				}

				pos := item - 49
				if counter[pos] > 0 {
					return false
				}
				counter[pos] = 1
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	isSudokuLine := checkLine(board)
	if !isSudokuLine {
		return false
	}
	isSudokuColumn := checkColumn(board)
	if !isSudokuColumn {
		return false
	}
	isSudokuBlock := checkBlock(board)
	if !isSudokuBlock {
		return false
	}
	return true
}
