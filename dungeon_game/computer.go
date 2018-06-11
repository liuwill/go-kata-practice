package dungeon_game

type MarkPoint struct {
	Energy int
	Best   int
}

func calculateKnightMinimumHP(dungeon [][]int) int {
	best := 0
	marks := make([][]MarkPoint, len((dungeon)))
	for index, line := range dungeon {
		marks[index] = make([]MarkPoint, len(line))
		for x := index; x < len(line); x++ {

		}

		if index < len(line)-1 {
			for y := index + 1; y < len(dungeon); y++ {

			}
		}
	}

	return best
}
