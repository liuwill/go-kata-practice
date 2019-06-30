package weekly_contest

func carPooling(trips [][]int, capacity int) bool {
	posData := make([]int, 1000)

	for pos := 0; pos < len(trips); pos++ {
		tripStep := trips[pos]
		num := tripStep[0]
		for i := tripStep[1]; i < tripStep[2]; i++ {
			posData[i-1] += num
		}
	}

	for num := range posData {
		if num > capacity {
			return false
		}
	}
	return true
}
