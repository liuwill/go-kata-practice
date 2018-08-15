package weekly_contest_96

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	times := 0
	sort.Sort(sort.Reverse(sort.IntSlice(people)))

	boat := 0
	index := 0

	// fmt.Printf("=== %v %v\n", people, limit)
	for index < len(people)-1 {
		for i := index; i < len(people); i++ {
			// fmt.Printf("%v %v %v %v\n", i, boat, index, limit)
			if boat == 0 && people[i] != 0 {
				boat += people[i]
				people[i] = 0
				index = i + 1
				continue
			}

			if people[i] != 0 && boat > 0 && boat+people[i] <= limit {
				boat += people[i]
				people[i] = 0
			}
		}

		if boat > 0 {
			times++
		}
		boat = 0
	}
	return times
}
