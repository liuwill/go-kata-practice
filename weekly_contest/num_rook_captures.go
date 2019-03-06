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

	for x := posY - 1; x >= 0; x-- {
		if board[x][posY] == ENEMY {
			count++
			break
		} else if board[x][posY] == FRIEND {
			break
		}
	}

	for x := posY + 1; x < len(board); x++ {
		if board[x][posY] == ENEMY {
			count++
			break
		} else if board[x][posY] == FRIEND {
			break
		}
	}

	return count
}
