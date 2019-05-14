package weekly_contest

func findFlower(source []int) int {
	for i := 0; i < 4; i++ {
		if source[i] == 0 {
			return i + 1
		}
	}
	return 1
}

/**
 * weekly-contest-1042
 * PUZZLE: Flower Planting With No Adjacent
 */
func gardenNoAdj(N int, paths [][]int) []int {
	sourceMap := make([][]int, N)
	for pos := 0; pos < len(paths); pos++ {
		matcher := paths[pos]
		i := matcher[0] - 1
		j := matcher[1] - 1
		sourceMap[i] = append(sourceMap[i], j)
		sourceMap[j] = append(sourceMap[j], i)
	}

	result := make([]int, N)
	result[0] = 1
	for i := 1; i < N; i++ {
		source := make([]int, 4)
		for j := 0; j < len(sourceMap[i]); j++ {
			pos := sourceMap[i][j]
			if result[pos] > 0 {
				source[result[pos]-1] = 1
			}
		}
		current := findFlower(source)
		result[i] = current
	}

	return result
}
