package weekly_contest_97

func superEggDrop(K int, N int) int {
	times := 0
	for N > 0 {
		times++
		N = N / 2
	}

	return times
}
