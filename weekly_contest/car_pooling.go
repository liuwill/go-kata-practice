package weekly_contest

/**
 * daily-challenge-1094
 * PUZZLE: Car Pooling
 */
func carPooling(trips [][]int, capacity int) bool {
	tripLen := len(trips)
	roadLen := 0
	for i := 0; i < len(trips); i++ {
		if trips[i][2] > roadLen {
			roadLen = trips[i][2]
		}
	}
	posData := make([]int, roadLen)

	for pos := 0; pos < tripLen; pos++ {
		tripStep := trips[pos]
		num := tripStep[0]
		for i := tripStep[1]; i < tripStep[2]; i++ {
			posData[i-1] += num
		}
	}

	for i := 0; i < len(posData); i++ {
		num := posData[i]
		if num > capacity {
			return false
		}
	}
	return true
}
