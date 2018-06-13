package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	length := len(gas)
	petrol := 0
	sum := 0

	for i := 0; i < length; i++ {
		sum += gas[i] - cost[i]
		petrol += gas[i] - cost[i]

		if petrol < 0 {
			petrol = 0
			start = i + 1
		}
	}

	if sum < 0 {
		start = -1
	}

	return start
}
