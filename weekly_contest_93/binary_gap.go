package weekly_contest_93

func binaryGap(N int) int {
	num := N
	result := 0
	length := 0
	for num > 0 {
		current := num % 2
		num = num / 2
		if current == 1 {
			length = 0
		} else {
			length++
		}

		if length > result {
			result = length
		}
	}

	return result
}
