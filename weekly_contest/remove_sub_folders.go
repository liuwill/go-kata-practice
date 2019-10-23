package weekly_contest

import "sort"

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

	if long[len(short)] != '/' {
		return 0
	}

	if position {
		return -1
	}

	return 1
}

/**
 * daily-challenge-1233
 * PUZZLE: Remove Sub-Folders from the Filesystem
 */
func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool { return len(folder[i]) < len(folder[j]) })

	result := []string{}
	rootFolder := make(map[string]bool)
	for _, path := range folder {
		isSub := false
		for i := 1; i < len(path); i++ {
			if path[i] == '/' && rootFolder[path[:i]] {
				isSub = true
				break
			}
		}

		if !isSub {
			rootFolder[path] = true
			result = append(result, path)
		}
	}

	return result
}

func removeSubfoldersSlow(folder []string) []string {
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
