package weekly_contest

import (
	"sort"
	"strings"
)

/**
 * daily-challenge-1268
 * PUZZLE: Search Suggestions System
 */
func suggestedProducts(products []string, searchWord string) [][]string {
	ll := len(searchWord)

	target := make([][]string, ll)
	left := products[:]
	sort.Strings(left)
	for i, _ := range searchWord {
		current := searchWord[0 : i+1]
		filtered := []string{}
		for _, item := range left {
			if strings.HasPrefix(item, current) {
				filtered = append(filtered, item)
			}
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
