package weekly_contest

/**
 * daily-challenge-1252
 * PUZZLE: Cells with Odd Values in a Matrix
 */
func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, m)
	}

	for _, v := range indices {
		for i := 0; i < m; i++ {
			matrix[v[0]][i]++
		}

		for j := 0; j < n; j++ {
			matrix[j][v[1]]++
		}
	}

	count := 0
	for _, line := range matrix {
		for _, item := range line {
			if item%2 == 1 {
				count++
			}
		}
	}

	return count
}
