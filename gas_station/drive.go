package gas_station

func next(current int, length int) int {
	return (current + 1) % length
}

func drive(start int, gas []int, cost []int) bool {
	length := len(gas)
	for i := 0; i < length; i++ {

	}
	return false
}

func canCompleteCircuit(gas []int, cost []int) int {
	origin := -1
	for i, _ := range gas {
		if drive(i, gas, cost) {
			origin = i
			break
		}
	}

	return origin
}
