package weekly_contest_97

func superEggDrop(K int, N int) int {
	times := 0

	tail := N
	left := K

	for left > 1 && tail > 2 {
		tail = tail / 2

		times++
		left--
	}

	if left >= 1 && tail > 0 {
		times = times + tail
		if N%2 == 0 && tail < N && tail > 2 {
			times--
		}
	}

	return times
}
