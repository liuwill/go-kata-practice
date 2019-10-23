package weekly_contest

func compareSubFolder(first string, second string) int {
	long := first
	short := second
	position := true

	if len(second) > len(first) {
		long = second
		short = first

		position = false
	}

	for i, _ := range short {
		if long[i] != short[i] {
			return 0
		}
	}

	if position {
		return -1
	}

	return 1
}

func removeSubfolders(folder []string) []string {
	mark := make([]int, len(folder))
	for i := 0; i < len(folder)-1; i++ {
		if mark[i] == 1 {
			continue
		}
		for j := i + 1; j < len(folder); j++ {
			if mark[j] == 1 {
				continue
			}

			compare := compareSubFolder(folder[i], folder[j])
			if compare == -1 {
				mark[i] = 1
			} else if compare == 1 {
				mark[j] = 1
			}
		}
	}

	result := []string{}
	for i, v := range mark {
		if v == 0 {
			result = append(result, folder[i])
		}
	}
	return result
}
