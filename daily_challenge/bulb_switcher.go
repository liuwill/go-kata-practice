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

func bulbSwitchEasy(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		pos := i
		round := 0
		for j := 1; j <= pos; j++ {
			if pos%j == 0 {
				round = round ^ 1
			}
		}

		if round == 1 {
			count++
		}
	}

	return count
}
