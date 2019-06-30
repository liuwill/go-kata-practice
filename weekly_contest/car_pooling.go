package weekly_contest

/**
 * daily-challenge-1094
 * PUZZLE: Car Pooling
 */
func carPooling(trips [][]int, capacity int) bool {
	tripLen := len(trips)
	roadLen := trips[tripLen-1][2]
	posData := make([]int, roadLen)

	for pos := 0; pos < len(trips); pos++ {
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
