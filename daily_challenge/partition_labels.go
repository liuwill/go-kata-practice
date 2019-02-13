package daily_challenge

/**
 * daily-challenge-763
 * PUZZLE: Partition Labels
 */
func partitionLabels(S string) []int {
	target := []int{}
	start := 0
	end := 0

	total := 0
	for i := 0; i < len(S); i++ {
		current := S[i]
		if i > end || i == len(S)-1 {
			// println(i, start, end)
			distance := end - start + 1
			target = append(target, distance)
			total += distance
			start = i
			end = i
		}

		for j := len(S) - 1; j > end; j-- {
			if current == S[j] {
				end = j
				break
			}
		}
	}

	if total < len(S) {
		target = append(target, len(S)-total)
	}
	return target
}
