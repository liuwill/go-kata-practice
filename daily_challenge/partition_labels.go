package daily_challenge

func partitionLabels(S string) []int {
	target := []int{}
	start := 0
	end := 0
	for i := 0; i < len(S); i++ {
		current := S[i]
		if i > end || i == len(S)-1 {
			// println(i, start, end)
			target = append(target, end-start+1)
			start = i
		}

		for j := len(S) - 1; j > i; j-- {
			if current == S[j] {
				end = j
				break
			}
		}
	}
	return target
}
