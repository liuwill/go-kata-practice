package island_number

import "fmt"

type Position struct {
	x int
	y int
}

func isMark(value byte) bool {
	if value == 0 || value == 2 {
		return true
	}
	return false
}

func scanIsland(grid [][]byte, positionList []Position) {
	fmt.Printf("%v", positionList)
	for {
		if len(positionList) < 1 {
			break
		}

		position := positionList[0]
		positionList = positionList[1:]

		for left := position.x - 1; left >= 0; left-- {
			if isMark(grid[left][position.y]) {
				break
			}

			grid[left][position.y] = 2
			positionList = append(positionList[:], Position{
				x: left,
				y: position.y,
			})
		}
		for top := position.y - 1; top >= 0; top-- {
			if isMark(grid[position.x][top]) {
				break
			}

			grid[position.x][top] = 2
			positionList = append(positionList[:], Position{
				x: position.x,
				y: top,
			})
		}
		for right := position.x + 1; right < len(grid[position.x]); right++ {
			if isMark(grid[right][position.y]) {
				break
			}

			grid[right][position.y] = 2
			positionList = append(positionList[:], Position{
				x: right,
				y: position.y,
			})
		}
		for bottom := position.y + 1; bottom < len(grid); bottom++ {
			if isMark(grid[position.x][bottom]) {
				break
			}

			grid[position.x][bottom] = 2
			positionList = append(positionList[:], Position{
				x: position.x,
				y: bottom,
			})
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
