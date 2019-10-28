package weekly_contest

func compareStraightLine(start []int, middle []int, end []int) bool {
	firstX := start[0] - middle[0]
	firstY := start[1] - middle[1]

	secondX := middle[0] - end[0]
	secondY := middle[1] - end[1]
	if (firstX == 0 || firstY == 0) && firstY == firstX {
		return true
	}
	if (secondX == 0 || secondY == 0) && secondY == secondX {
		return true
	}

	if firstY != 0 && secondY != 0 {
		if firstX/firstY == secondX/secondY {
			return true
		}
	}

	return false
}

func checkStraightLine(coordinates [][]int) bool {
	for i := 1; i < len(coordinates)-1; i++ {
		if !compareStraightLine(coordinates[i-1], coordinates[i], coordinates[i+1]) {
			return false
		}
	}
	return true
}
