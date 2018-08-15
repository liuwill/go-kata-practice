package weekly_contest_96

import (
	"sort"
)

func numRescueBoatsSort(people []int, limit int) int {
	// sort.Sort(sort.Reverse(sort.IntSlice(people)))
	sort.Ints(people)

	times := 0
	boat := 0
	index := len(people) - 1

	// fmt.Printf("=== %v %v\n", limit, len(people))
	for index > 0 {
		pos := index
		number := 0
		for i := index; i >= 0; i-- {
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
		index = pos - 1
		boat = 0
	}

	if index >= 0 && people[index] > 0 {
		times++
	}
	return times
}
