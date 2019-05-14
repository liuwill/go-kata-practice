package weekly_contest

import "sort"

func findFlower(source []int) int {
	sort.Ints(source)
	for i := 0; i < 3; i++ {
		if source[i] == 1 {
			return source[i] + 1
		}
	}
	return 1
}

func gardenNoAdj(N int, paths [][]int) []int {
	sourceMap := make([][]int, N)
	for pos := 0; pos < len(paths); pos++ {
		matcher := paths[pos]
		i := matcher[0]
		j := matcher[1]
		sourceMap[i] = append(sourceMap[i], j)
		sourceMap[j] = append(sourceMap[j], i)
	}

	result := make([]int, N)
	result[0] = 1
	for i := 1; i < N; i++ {
		source := make([]int, 4)
		for j := 0; j < len(sourceMap[i]); j++ {
			pos := sourceMap[i][j]
			source[result[pos]-1] = 1
		}
		current := findFlower(source)
		result[i] = current
	}

	return result
}
