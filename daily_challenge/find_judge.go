package daily_challenge

func findJudge(N int, trust [][]int) int {
	trustMarks := make([]int, N)
	trustedMarks := make([]int, N)

	for _, relation := range trust {
		trustMarks[relation[0]-1] = 1
		trustedMarks[relation[1]-1] += 1
	}

	target := -1
	for i := 1; i <= N; i++ {
		pos := i - 1
		if trustMarks[pos] == 0 && trustedMarks[pos] == N-1 {
			if target > 0 {
				return -1
			}
			target = i
		}
	}
	return target
}
