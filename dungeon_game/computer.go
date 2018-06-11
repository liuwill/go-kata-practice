package dungeon_game

type MarkPoint struct {
	Energy int
	Best   int
}

func calculateCurrentMark(x int, y int, dungeon [][]int, marks [][]MarkPoint) MarkPoint {
	energy := 0
	best := 0

	if x == 0 && y == 0 {
		energy = 1 - dungeon[0][0]
		best = 1
		if energy > 1 {
			best = energy
		}
	} else {
		if x > 0 {
			energy = marks[x-1][y].Energy
			best = marks[x-1][y].Best
		}
		if y > 0 && marks[x][y-1].Best < best {
			energy = marks[x][y-1].Energy
			best = marks[x][y-1].Best
		}
	}

	return MarkPoint{
		Energy: energy,
		Best:   best,
	}
}

func calculateMinimumHP(dungeon [][]int) int {
	best := 0
	width := 0
	height := len(dungeon)

	marks := make([][]MarkPoint, len(dungeon))
	for index, line := range dungeon {
		marks[index] = make([]MarkPoint, len(line))
		width = len(line)
		for x := index; x < len(line); x++ {
			marks[index][x] = calculateCurrentMark(index, x, dungeon, marks)
		}

		if index < len(line)-1 {
			for y := index + 1; y < len(dungeon); y++ {
				// println(y, index)
				marks[y][index] = calculateCurrentMark(y, index, dungeon, marks)
			}
		}
	}

	best = marks[height-1][width-1].Best
	return best
}
