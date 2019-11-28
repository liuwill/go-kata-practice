package weekly_contest

import "strings"

func suggestedProducts(products []string, searchWord string) [][]string {
	ll := len(searchWord)

	target := make([][]string, ll)
	left := make([]string, len(products))
	for i, v := range products {
		left[i] = v
	}
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
