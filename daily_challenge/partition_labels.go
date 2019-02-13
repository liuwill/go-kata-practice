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

/**
 * 在检查之前，可以通过一次循环。找到每个字符所在的最右边的位置，避免在检查循环中内嵌for循环
 */
func partitionLabelsFast(S string) []int {
	target := []int{}
	start := 0
	end := 0

	total := 0
	posMap := [26]int{}
	for i := range posMap {
		posMap[i] = -1
	}
	for i, val := range S {
		pos := val - 'a'
		posMap[pos] = i
	}

	for i := 0; i < len(S); i++ {
		current := S[i]
		pos := current - 'a'

		if i > end || i == len(S)-1 {
			// println(i, start, end)
			distance := end - start + 1
			target = append(target, distance)
			total += distance
			start = i
			end = i
		}

		if posMap[pos] > end {
			end = posMap[pos]
		}
	}

	if total < len(S) {
		target = append(target, len(S)-total)
	}
	return target
}
