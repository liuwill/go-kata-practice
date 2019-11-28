package weekly_contest

import (
	"sort"
)

/**
 * daily-challenge-1268
 * PUZZLE: Search Suggestions System
 */
func suggestedProducts(products []string, searchWord string) [][]string {
	ll := len(searchWord)

	target := make([][]string, ll)
	left := products[:]
	sorted := false
	for i, _ := range searchWord {
		wl := i + 1
		current := searchWord[:wl]
		filtered := []string{}
		for _, word := range left {
			if i < len(word) && word[:wl] == current {
				filtered = append(filtered, word)
			}
		}
		if !sorted {
			sort.Strings(filtered)
			sorted = true
		}

		left = filtered

		l := len(filtered)
		if l > 3 {
			l = 3
		}

		target[i] = filtered[0:l]
	}

	return target
}
