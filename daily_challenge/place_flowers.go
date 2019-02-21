package daily_challenge

/**
 * daily-challenge-605
 * PUZZLE: Can Place Flowers
 */
func canPlaceFlowers(flowerbed []int, n int) bool {
	previous := 0

	for i := 0; i < len(flowerbed)-1; i++ {
		current := flowerbed[i]
		next := 0
		if i+1 < len(flowerbed) {
			next = flowerbed[i+1]
		}

		if previous == 0 && current == 0 && next == 0 {
			flowerbed[i] = 1
			n--
		}

		previous = flowerbed[i]
	}
	return n <= 0
}
