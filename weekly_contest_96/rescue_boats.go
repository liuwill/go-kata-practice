package weekly_contest_96

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	i := 0
	j := len(people) - 1
	times := 0

	for i <= j {
		times++
		if people[i]+people[j] <= limit {
			i++
		}

		j--
	}

	return times
}
