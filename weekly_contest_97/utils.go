package weekly_contest_97

func compareListWithoutOrder(expect []string, target []string) bool {
	if len(expect) != len(target) {
		return false
	}

	count := 0
	for _, exp := range expect {
		for _, tar := range target {
			if exp == tar {
				count++
				break
			}
		}
	}
	return count == len(expect)
}
