package weekly_contest

func tanStraightLine(start []int, end []int) int {
	return (start[0] - end[0]) / (start[1] - end[1])
}

func checkStraightLine(coordinates [][]int) bool {
	tan := tanStraightLine(coordinates[0], coordinates[1])
	for i := 1; i < len(coordinates)-1; i++ {
		current := tanStraightLine(coordinates[i], coordinates[i+1])
		if tan != current {
			return false
		}
	}
	return true
}
