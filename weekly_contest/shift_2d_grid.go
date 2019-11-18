package weekly_contest

func findShiftPosition(hl int, wl int, x int, y int, round int) []int {
	yd := round / wl
	xd := round % wl

	xp := x - xd
	if yd > hl {
		yd = yd % hl
	}

	return []int{xp, yd}
}

func shiftGrid(grid [][]int, k int) [][]int {
	hl := len(grid)
	wl := len(grid[0])
	target := make([][]int, hl)
	for i := 0; i < hl; i++ {
		target[i] = make([]int, wl)

		for j := 0; j < wl; j++ {
			p := findShiftPosition(hl-1, wl-1, i, j, k)
			target[i][j] = grid[p[0]][p[1]]
		}
	}

	return target
}
