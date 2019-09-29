package weekly_contest

func uniqueOccurrences(arr []int) bool {
	result := make([]int, 1001)
	dict := make(map[int]int)

	for _, v := range arr {
		if _, ok := dict[v]; !ok {
			dict[v] = 0
		}
		dict[v]++
	}
	for _, v := range dict {
		if result[v] > 0 {
			return false
		}
	}

	return true
}
