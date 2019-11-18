package weekly_contest

func findShiftPosition(hl int, wl int, y int, x int, round int) []int {
	yd := round / wl
	xd := round % wl

	xp := x - xd
	if xp < 0 {
		xp = wl + xp
		yd++
	}

	if yd > hl {
		yd = yd % hl
	}
	yp := y - yd
	if yp < 0 {
		yp = hl + yp
	}

	return []int{yp, xp}
}

func shiftGrid(grid [][]int, k int) [][]int {
	hl := len(grid)
	wl := len(grid[0])
	target := make([][]int, hl)
	for i := 0; i < hl; i++ {
		target[i] = make([]int, wl)

		for j := 0; j < wl; j++ {
			p := findShiftPosition(hl, wl, i, j, k)
			target[i][j] = grid[p[0]][p[1]]
		}
	}

	return target
}
