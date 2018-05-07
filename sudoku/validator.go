package sudoku

func checkLine(board [][]byte) bool {
	for _, line := range board {
		if len(line) != 9 {
			return false
		}

	}
	return true
}

func checkColumn(board [][]byte) bool { return false }

func checkBlock(board [][]byte) bool { return false }

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
