package daily_challenge

/**
 * daily-challenge-605
 * PUZZLE: Can Place Flowers
 */
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 1; i < len(flowerbed); i++ {
		previous := flowerbed[i-1]
		current := flowerbed[i]
		next := 0
		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}

		if previous == 0 && current == 0 && next == 0 {
			n--
		}
	}
	return n == 0
}
