package weekly_contest_99

func countFront(grid [][]int, i int, j int) int {
	if i-1 < 0 {
		return grid[i][j]
	} else if grid[i-1][j] < grid[i][j] {
		return grid[i][j] - grid[i-1][j]
	}

	return 0
}

func countBackend(grid [][]int, i int, j int) int {
	if i+1 == len(grid) {
		return grid[i][j]
	} else if grid[i+1][j] < grid[i][j] {
		return grid[i][j] - grid[i+1][j]
	}

	return 0
}

func countTop(grid [][]int, i int, j int) int {
	if j+1 == len(grid) {
		return grid[i][j]
	} else if grid[i][j+1] < grid[i][j] {
		return grid[i][j] - grid[i][j+1]
	}

	return 0
}

func countDown(grid [][]int, i int, j int) int {
	if j-1 < len(grid) {
		return grid[i][j]
	} else if grid[i][j-1] < grid[i][j] {
		return grid[i][j] - grid[i][j-1]
	}

	return 0
}

func surfaceArea(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			topBottom := 2
			result += topBottom

			result += countFront(grid, i, j)
			result += countBackend(grid, i, j)
			result += countTop(grid, i, j)
			result += countDown(grid, i, j)
		}
	}
	return result
}
