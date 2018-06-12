package dungeon_game

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func countMarks(i int, j int, dungeon [][]int, marks [][]int) int {
	width := len(dungeon[0])
	height := len(dungeon)

	mark := marks[i][j]
	if i+1 <= height-1 && j+1 <= width-1 {
		mark = min(marks[i+1][j], marks[i][j+1]) - dungeon[i][j]
	} else if i+1 <= height-1 {
		mark = marks[i+1][j] - dungeon[i][j]
	} else if j+1 <= width-1 {
		mark = marks[i][j+1] - dungeon[i][j]
	}

	if mark <= 0 {
		mark = 1
	}
	return mark
}

func calculateMinimumHP(dungeon [][]int) int {
	width := len(dungeon[0])
	height := len(dungeon)

	marks := make([][]int, len(dungeon))
	for index, line := range dungeon {
		marks[index] = make([]int, len(line))
	}

	marks[height-1][width-1] = max(-dungeon[height-1][width-1], 0) + 1
	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			marks[i][j] = countMarks(i, j, dungeon, marks)
		}
	}

	return marks[0][0]
}

/*
func calculateMinimumHP(dungeon [][]int) int {
	width := len(dungeon[0])
	height := len(dungeon)

	marks := make([][]int, len(dungeon))
	for index, line := range dungeon {
		marks[index] = make([]int, len(line))
	}

	marks[height-1][width-1] = max(-dungeon[height-1][width-1], 0) + 1
	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i+1 <= height-1 && j+1 <= width-1 {
        marks[i][j] = min(marks[i+1][j], marks[i][j+1]) - dungeon[i][j]
      } else if i+1 <= height-1 {
        marks[i][j] = marks[i+1][j] - dungeon[i][j]
      } else if j+1 <= width-1 {
        marks[i][j] = marks[i][j+1] - dungeon[i][j]
      }

      if marks[i][j] <= 0 {
        marks[i][j] = 1
      }
		}
	}

	return marks[0][0]
}
*/
