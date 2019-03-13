package weekly_contest

import "math"

func powerInt(num int, power int) int {
	result := math.Pow(float64(num), float64(power))
	return int(result)
}

/**
 * daily-challenge-970
 * PUZZLE: Powerful Integers
 */
func powerfulIntegers(x int, y int, bound int) []int {
	source := make([]int, bound+1)
	for i := 0; powerInt(x, i)+1 <= bound; i++ {
		for j := 0; powerInt(y, j)+1 <= bound; j++ {
			val := powerInt(x, i) + powerInt(y, j)
			if val <= bound {
				source[val]++
			}
		}
	}

	target := []int{}
	for pos, val := range source {
		if val > 0 {
			target = append(target, pos)
		}
	}
	return target
}
