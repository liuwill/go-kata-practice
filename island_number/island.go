package island_number

const ISLAND_MARK byte = 2

type Position struct {
	x int
	y int
}

var moveActions = []func(x int, y int) Position{
	func(x int, y int) Position {
		return Position{
			x: x - 1,
			y: y,
		}
	},
	func(x int, y int) Position {
		return Position{
			x: x + 1,
			y: y,
		}
	},
	func(x int, y int) Position {
		return Position{
			x: x,
			y: y - 1,
		}
	},
	func(x int, y int) Position {
		return Position{
			x: x,
			y: y + 1,
		}
	},
}

func isVirginIsland(value byte) bool {
	if value == '1' || value == 1 {
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

func scanIsland(grid [][]byte, position Position) {
	positionList := []Position{
		position,
	}

	for len(positionList) > 0 {
		position := positionList[0]
		positionList = positionList[1:]

		for _, doAction := range moveActions {
			nextPos := doAction(position.x, position.y)
			if !isInMap(grid, nextPos) {
				continue
			}
			if !isVirginIsland(grid[nextPos.x][nextPos.y]) {
				continue
			}

			grid[nextPos.x][nextPos.y] = ISLAND_MARK
			positionList = append(positionList[:], Position{
				x: nextPos.x,
				y: nextPos.y,
			})
		}
	}
}

func numIslands(grid [][]byte) int {
	islandNumbers := 0

	for i := 0; i < len(grid); i++ {
		column := grid[i]
		for j := 0; j < len(column); j++ {
			if !isVirginIsland(grid[i][j]) {
				continue
			}

			islandNumbers++
			grid[i][j] = ISLAND_MARK
			scanIsland(grid[:], Position{
				x: i,
				y: j,
			})
		}
	}
	return islandNumbers
}
