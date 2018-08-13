package weekly_contest_97

func superEggDrop(K int, N int) int {
	times := 0

	tail := N
	left := K

	for K > 1 && tail > 0 {
		tail = tail / 2
		times++
		left--
	}

	if K == 1 && tail > 0 {
		times = times + tail
	}

	return times
}
