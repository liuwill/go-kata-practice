package weekly_contest_93

func binaryGap(N int) int {
	// println("------", N)
	num := N
	result := 0
	length := 0
	start := 0
	for num > 0 {
		current := num % 2
		num = num / 2

		// println(current, length, start, result)
		if current == 1 {
			if length+1 > result && start == 1 {
				result = length + 1
				start = 0
			}
			start = 1
			length = 0
		} else if current == 0 && start == 1 {
			length++
		}
	}

	return result
}
