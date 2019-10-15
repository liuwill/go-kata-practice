package weekly_contest

const MAP_SIZE = 8

func positionAttackQueen(x int, y int) int {
	if x < 0 && y == 0 {
		return 0
	} else if x < 0 && y < 0 {
		return 1
	} else if x == 0 && y < 0 {
		return 2
	} else if x > 0 && y < 0 {
		return 3
	} else if x > 0 && y == 0 {
		return 4
	} else if x < 0 && y > 0 {
		return 5
	} else if x == 0 && y > 0 {
		return 6
	} // else if x < 0 0 && y > 0 { return 7 }

	return 7
}

/**
 * daily-challenge-1222
 * PUZZLE: Queens That Can Attack the King
 */
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	direction := make([][]int, MAP_SIZE)
	for i, _ := range direction {
		direction[i] = []int{-1, 100}
	}

	count := 0
	for i, q := range queens {
		x := q[0] - king[0]
		y := q[1] - king[1]

		xx := x * x
		yy := y * y

		distance := xx + yy
		if xx == yy || x == 0 || y == 0 {
			pos := positionAttackQueen(x, y)
			if direction[pos][1] > distance {
				if direction[pos][0] < 0 {
					count++
				}
				direction[pos] = []int{i, distance}
			}
		}
	}

	result := make([][]int, count)
	p := 0
	for _, d := range direction {
		if d[0] != -1 {
			result[p] = queens[d[0]]
			p++
		}
	}

	return result
}
