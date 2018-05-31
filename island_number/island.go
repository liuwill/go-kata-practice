package island_number

type Position struct {
	x int
	y int
}

var moveActions = []func(grid [][]byte, x int, y int) Position{
	func(grid [][]byte, x int, y int) Position {
		return Position{
			x: x - 1,
			y: y,
		}
	},
	func(grid [][]byte, x int, y int) Position {
		return Position{
			x: x + 1,
			y: y,
		}
	},
	func(grid [][]byte, x int, y int) Position {
		return Position{
			x: x,
			y: y - 1,
		}
	},
	func(grid [][]byte, x int, y int) Position {
		return Position{
			x: x,
			y: y + 1,
		}
	},
}

func isMark(value byte) bool {
	if value == '0' || value == 0 || value == 2 {
		return true
	}
	return false
}

func isInMap(grid [][]byte, pos Position) bool {
	x := pos.x
	y := pos.y
	if x < 0 || y < 0 {
		return false
	}

	if x >= len(grid) || y >= len(grid[x]) {
		return false
	}

	return true
}

func scanIsland(grid [][]byte, positionList []Position) {
	for {
		if len(positionList) < 1 {
			break
		}

		position := positionList[0]
		positionList = positionList[1:]

		for _, action := range moveActions {
			nextPos := action(grid, position.x, position.y)
			if !isInMap(grid, nextPos) {
				continue
			}
			if !isMark(grid[nextPos.x][nextPos.y]) {
				grid[nextPos.x][nextPos.y] = 2
				positionList = append(positionList[:], Position{
					x: nextPos.x,
					y: nextPos.y,
				})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	current := 0
	// var data [][]byte
	// copy(data, grid)
	positionList := []Position{}

	for i := 0; i < len(grid); i++ {
		column := grid[i]
		for j := 0; j < len(column); j++ {
			mark := grid[i][j]
			if isMark(mark) {
				continue
			}

			current++
			grid[i][j] = 2
			positionList = append(positionList[:], Position{
				x: i,
				y: j,
			})
			scanIsland(grid[:], positionList[:])
		}
	}
	return current
}
