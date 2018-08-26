package weekly_contest_93

// target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]

// target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]

// target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,0]]
// target = 100, startFuel = 10, stations = [[10,60],[20,20],[30,30],[60,10]]

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	current := 0
	fuel := startFuel
	index := 0
	finish := false
	for current < target && fuel > 0 && !finish {
		left := fuel
		// maxLeft := 0
		// maxIndex := 0
		for i := index; i < len(stations); i++ {
			distance := stations[i][0]
			supplier := stations[i][1]

			left = left - (distance - current)
			if left < 0 {
				break
			}
			left = left + supplier
			if 100-distance <= left {
				finish = true
			}
		}

	}
	return index
}
