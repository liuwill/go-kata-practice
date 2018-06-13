package gas_station

func next(current int, length int) int {
	return (current + 1) % length
}

func drive(start int, gas []int, cost []int) bool {
	length := len(gas)
	petrol := 0
	current := start
	for i := 0; i < length; i++ {
		petrol += gas[current]

		if petrol >= cost[current] {
			petrol -= cost[current]
			current = next(current, length)
		} else {
			return false
		}
	}
	return true
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
