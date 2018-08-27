package weekly_contest_93

// target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]

// target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]

// target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,0]]
// target = 100, startFuel = 10, stations = [[10,60],[20,20],[30,30],[60,10]]

func MaxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	length := len(stations)
	list := make([]int, length+1) // new long[length + 1]
	list[0] = startFuel
	for i := 0; i < length; i++ {
		for t := i; t >= 0; t-- {
			if list[t] >= stations[i][0] {
				list[t+1] = MaxInt(list[t+1], list[t]+stations[i][1])
			}
		}
	}

	for i := 0; i <= length; i++ {
		if list[i] >= target {
			return i
		}
	}
	return -1
}
