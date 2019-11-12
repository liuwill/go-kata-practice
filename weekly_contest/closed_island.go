package weekly_contest

const (
	ISLAND_POINT_TYPE_WATER   = 1
	ISLAND_POINT_TYPE_ISLAND  = 0
	ISLAND_POINT_TYPE_DANGER  = 2
	ISLAND_POINT_TYPE_VISITED = 3
)

var closedMoveActions = []func(x int, y int) []int{
	func(x int, y int) []int {
		return []int{x - 1, y}
	},
	func(x int, y int) []int {
		return []int{x + 1, y}
	},
	func(x int, y int) []int {
		return []int{x, y - 1}
	},
	func(x int, y int) []int {
		return []int{x, y + 1}
	},
}

func isInClosedMap(grid [][]int, pos []int) bool {
	x := pos[0]
	y := pos[1]
	if x < 0 || y < 0 {
		return false
	}

	if x >= len(grid) || y >= len(grid[x]) {
		return false
	}

	return true
}

func scanClosedIsland(grid [][]int, point []int) bool {
	isMatch := true
	positionList := [][]int{point}

	for len(positionList) > 0 {
		position := positionList[0]
		positionList = positionList[1:]

		for _, doAction := range closedMoveActions {
			nextPos := doAction(position[0], position[1])
			if !isInClosedMap(grid, nextPos) {
				isMatch = false
				continue
			}
			if ISLAND_POINT_TYPE_ISLAND != grid[nextPos[0]][nextPos[1]] {
				continue
			}

			grid[nextPos[0]][nextPos[1]] = ISLAND_POINT_TYPE_VISITED
			positionList = append(positionList[:], nextPos)
		}
	}

	return isMatch
}

func closedIsland(grid [][]int) int {
	islandNumbers := 0

	for i, column := range grid {
		for j, _ := range column {
			if ISLAND_POINT_TYPE_ISLAND != grid[i][j] {
				continue
			}

			grid[i][j] = ISLAND_POINT_TYPE_VISITED
			if scanClosedIsland(grid[:], []int{i, j}) {
				islandNumbers++
			}
		}
	}

	return islandNumbers
}
