package weekly_contest

/**
 * daily-challenge-999
 * PUZZLE: Available Captures for Rook
 */

const (
	CURRENT = 'R'
	FRIEND  = 'B'
	ENEMY   = 'p'
	POINT   = '.'
)

func numRookCaptures(board [][]byte) int {
	count := 0
	posX := -1
	posY := -1
	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == CURRENT {
				posX = i
				posY = j
				break
			}
		}
	}

	for y := posY - 1; y >= 0; y-- {
		if board[posX][y] == ENEMY {
			count++
			break
		} else if board[posX][y] == FRIEND {
			break
		}
	}

	for y := posY + 1; y < len(board[posX]); y++ {
		if board[posX][y] == ENEMY {
			count++
			break
		} else if board[posX][y] == FRIEND {
			break
		}
	}

	for x := posX - 1; x >= 0; x-- {
		if board[x][posY] == ENEMY {
			count++
			break
		} else if board[x][posY] == FRIEND {
			break
		}
	}

	for x := posX + 1; x < len(board); x++ {
		if board[x][posY] == ENEMY {
			count++
			break
		} else if board[x][posY] == FRIEND {
			break
		}
	}

	return count
}

/*
func numRookCapturesFast(board [][]byte) int {
	count := 0
	for i, _ := range board {
		mode := false
		scanned := false
		rowCount := 0

		for j, _ := range board[i] {
			if board[i][j] == CURRENT {
				mode = true
				if scanned {
					rowCount++
				}
			}

			if mode {
				if board[i][j] == ENEMY {
					rowCount++
				}
				if board[i][j] != POINT {
					break
				}
			} else {
				if board[i][j] == ENEMY {
					scanned = true
				} else if board[i][j] == POINT {
					scanned = false
				}
			}
		}

		count += rowCount
	}

	return count
}
*/
