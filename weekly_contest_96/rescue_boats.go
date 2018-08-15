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
		pos := index
		number := 0
		for i := index; i < len(people); i++ {
			// fmt.Printf("i=%v boat=%v index=%v %v -> %v\n", i, boat, index, limit, times)
			if people[i] == 0 {
				continue
			}

			if boat == 0 && people[i] != 0 {
				boat += people[i]
				people[i] = 0
				pos = i
				number = 1
			}

			if people[i] != 0 && boat > 0 && boat+people[i] <= limit {
				boat += people[i]
				people[i] = 0
				number++
			}

			if boat == limit || number > 1 {
				break
			}
		}

		if boat > 0 {
			times++
		}
		index = pos + 1
		boat = 0
	}

	if index < len(people) && people[index] > 0 {
		times++
	}
	return times
}
