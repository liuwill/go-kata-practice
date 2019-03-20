package daily_challenge

/**
 * daily-challenge-319
 * PUZZLE: Bulb Switcher
 */
func bulbSwitch(n int) int {
	list := make([]int, n)

	for i := 0; i < n; i++ {
		for pos := i; pos < n; pos += i + 1 {
			list[pos] = list[pos] ^ 1
		}
	}

	target := 0
	for _, val := range list {
		if val == 1 {
			target++
		}
	}
	// fmt.Printf("BULB: %v \n", list)
	return target
}
