package weekly_contest_97

func superEggDrop(K int, N int) int {
	/*
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
	*/
	lo := 1
	hi := N
	for lo < hi {
		mi := (lo + hi) / 2
		if countEgg(mi, K, N) < N {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return lo
}

func countEgg(x int, K int, N int) int {
	ans := 0
	r := 1
	for i := 1; i <= K; i++ {
		r = r * (x - i + 1)
		r /= i
		ans += r
		if ans >= N {
			break
		}
	}
	return ans
}
