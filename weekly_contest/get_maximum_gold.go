package weekly_contest

/**
 * weekly-contest-1219
 * PUZZLE: Path with Maximum Gold
 */
func deepGoldSearch(i int, j int, grid [][]int, target [][]int, visited [][]int, current int) {
	if i < 0 || i > len(grid)-1 ||
		j < 0 || j > len(grid[0])-1 ||
		grid[i][j] == 0 || visited[i][j] == 1 {
		return
	}
	visited[i][j] = 1
	if current > target[i][j] {
		target[i][j] = current
	}

	if i > 0 {
		deepGoldSearch(i-1, j, grid, target, visited, current+grid[i-1][j])
	}
	if j > 0 {
		deepGoldSearch(i, j-1, grid, target, visited, current+grid[i][j-1])
	}
	if i < len(grid)-1 {
		deepGoldSearch(i+1, j, grid, target, visited, current+grid[i+1][j])
	}
	if j < len(grid[0])-1 {
		deepGoldSearch(i, j+1, grid, target, visited, current+grid[i][j+1])
	}
	visited[i][j] = 0
}

func initGoldMap(grid [][]int) [][]int {
	target := make([][]int, len(grid))

	for i, line := range grid {
		target[i] = make([]int, len(line))
	}

	return target
}

func getMaximumGold(grid [][]int) int {
	visited := make([][]int, len(grid))
	target := make([][]int, len(grid))

	for i, line := range grid {
		visited[i] = make([]int, len(line))
		target[i] = make([]int, len(line))
	}

	for i, line := range grid {
		for j, item := range line {
			visited = initGoldMap(grid)
			deepGoldSearch(i, j, grid, target, visited, item)
		}
	}

	result := 0
	for _, line := range target {
		for _, item := range line {
			if item > result {
				result = item
			}
		}
	}
	return result
}
