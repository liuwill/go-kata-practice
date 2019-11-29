package weekly_contest

/**
 * daily-challenge-1267
 * PUZZLE: Count Servers that Communicate
 */
func countServers(grid [][]int) int {
	count := 0
	m := len(grid)
	n := len(grid[0])

	row := make([]int, m)
	column := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[i] += grid[i][j]
			column[j] += grid[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}

			if row[i] > 1 || column[j] > 1 {
				count++
			}
		}
	}

	return count
}
