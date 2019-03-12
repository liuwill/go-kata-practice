package weekly_contest

/**
 * daily-challenge-994
 * PUZZLE: Rotting Oranges
 */

const (
	ROTTING_ORANGE = 2
	FRESH_ORANGE   = 1
)

func fetchPoint(grid [][]int, x int, y int) int {
	if x < 0 || y < 0 {
		return -1
	} else if x >= len(grid) || y >= len(grid[0]) {
		return -1
	}

	return grid[x][y]
}

var actionMap = []func(x int, y int) []int{
	func(x int, y int) []int {
		return []int{x - 1, y}
	},
	func(x int, y int) []int {
		return []int{x, y - 1}
	},
	func(x int, y int) []int {
		return []int{x + 1, y}
	},
	func(x int, y int) []int {
		return []int{x, y + 1}
	},
}

func orangesRotting(grid [][]int) int {
	hasRotting := true
	round := 0
	for hasRotting {
		hasRotting = false
		rottingList := [][]int{}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] != ROTTING_ORANGE {
					continue
				}

				for _, actionFunc := range actionMap {
					position := actionFunc(i, j)
					val := fetchPoint(grid, position[0], position[1])
					if val == FRESH_ORANGE {
						rottingList = append(rottingList, position)
					}
				}
			}
		}

		if len(rottingList) > 0 {
			hasRotting = true
			round++
		}

		for _, position := range rottingList {
			x := position[0]
			y := position[1]
			grid[x][y] = ROTTING_ORANGE
		}
	}
	return len(grid)
}
