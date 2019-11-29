package weekly_contest

func countServers(grid [][]int) int {
	count := 0
	m := len(grid)
	n := len(grid[0])

	source := make([][]int, m)
	for i, line := range grid {
		source[i] = line[:]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if source[i][j] == 0 {
				continue
			}

			if i-1 >= 0 && source[i-1][j] > 0 {
				source[i][j]++
				count++
				continue
			}
			if j-1 >= 0 && source[i][j-1] > 0 {
				source[i][j]++
				count++
				continue
			}

			if i+1 < m && source[i+1][j] > 0 {
				source[i][j]++
				count++
				continue
			}

			if j+1 < n && source[i][j+1] > 0 {
				source[i][j]++
				count++
				continue
			}
		}
	}
	return count
}
