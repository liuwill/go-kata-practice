package weekly_contest

func farLandDistance(x []int, y []int) int {
	long := y[1] - x[1]
	width := y[0] - x[0]

	if long < 0 {
		long = -long
	}
	if width < 0 {
		width = -width
	}

	return long + width
}

/**
 * daily-challenge-1162
 * PUZZLE: As Far from Land as Possible
 */
func maxDistance(grid [][]int) int {
	max := 0
	N := len(grid)

	container := make([][]int, N)
	for i, _ := range container {
		container[i] = make([]int, N)
	}

	for i, line := range grid {
		for j, v := range line {
			if v == 0 {
				continue
			}

			for ii, line := range grid {
				for jj, v := range line {
					if v == 1 {
						continue
					}

					distance := farLandDistance([]int{i, j}, []int{ii, jj})
					if container[ii][jj] == 0 || distance < container[ii][jj] {
						container[ii][jj] = distance
					}
				}
			}
		}
	}

	for _, line := range container {
		for _, v := range line {
			if v > max {
				max = v
			}
		}
	}
	if max == 0 {
		max = -1
	}
	return max
}
