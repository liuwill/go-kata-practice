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

func maxDistance(grid [][]int) int {
	max := 0
	for i, line := range grid {
		for j, v := range line {
			if v == 1 {
				continue
			}

			for ii, line := range grid {
				for jj, v := range line {
					if v == 0 {
						continue
					}

					distance := farLandDistance([]int{i, j}, []int{ii, jj})
					if distance > max {
						max = distance
					}
				}
			}
		}
	}
	return max
}
