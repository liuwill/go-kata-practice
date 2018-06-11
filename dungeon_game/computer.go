package dungeon_game

type MarkPoint struct {
	Energy int
	Best   int
}

func calculateCurrentMark(x int, y int, dungeon [][]int, marks [][]MarkPoint) MarkPoint {
	energy := 0
	best := 0

	if x == 0 && y == 0 {
		energy = 1
		best = 1
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
	nextEnergy := energy - dungeon[x][y]
	currentBest := best
	if currentBest < nextEnergy {
		currentBest = nextEnergy
	}

	target := MarkPoint{
		Energy: nextEnergy,
		Best:   currentBest,
	}
	// fmt.Printf("=== %v => ", target)
	// println(x, y, energy, best)
	return target
}

func calculateMinimumHP(dungeon [][]int) int {
	best := 0
	width := 0
	height := len(dungeon)

	marks := make([][]MarkPoint, len(dungeon))
	for index, line := range dungeon {
		marks[index] = make([]MarkPoint, len(line))
	}
	for index, line := range dungeon {
		width = len(line)
		// fmt.Printf("%v\n", marks)
		for x := index; x < len(line); x++ {
			marks[index][x] = calculateCurrentMark(index, x, dungeon, marks)
			// println("++", index, x, dungeon[index][x])
		}

		if index < len(line)-1 {
			for y := index + 1; y < len(dungeon); y++ {
				marks[y][index] = calculateCurrentMark(y, index, dungeon, marks)
			}
		}
	}

	best = marks[height-1][width-1].Best
	return best
}
