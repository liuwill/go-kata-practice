package weekly_contest

const MAX_COORDINATES_POINT = 10001

func tanStraightLine(start []int, end []int) int {
	x := start[0] - end[0]
	y := start[1] - end[1]

	if x == 0 {
		return 0
	}
	if y == 0 {
		return MAX_COORDINATES_POINT
	}

	return x / y
}

/**
 * daily-challenge-1232
 * PUZZLE: Check If It Is a Straight Line
 */
func checkStraightLine(coordinates [][]int) bool {
	tan := tanStraightLine(coordinates[0], coordinates[1])
	for i := 1; i < len(coordinates)-1; i++ {
		current := tanStraightLine(coordinates[i], coordinates[i+1])
		if current != tan {
			return false
		}
	}
	return true
}

func checkStraightLineFast(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	base1 := coordinates[0]
	base2 := coordinates[1]

	sub0 := base2[0] - base1[0]
	sub1 := base2[1] - base1[1]
	for i := 2; i < len(coordinates); i++ {
		current := coordinates[i]
		if (current[1]-base2[1])*sub0 != sub1*(current[0]-base2[0]) {
			return false
		}
	}
	return true
}
