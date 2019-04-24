package weekly_contest

import (
	"sort"
)

type DistanceList [][]int

func (self DistanceList) Len() int {
	return len(self)
}

func (self DistanceList) Less(i, j int) bool {
	return self[i][1] > self[j][1]
}

func (self DistanceList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

/**
 * weekly-contest-1029
 * PUZZLE: Two City Scheduling
 */
func twoCitySchedCost(costs [][]int) int {
	target := 0
	disList := make([][]int, len(costs))
	count := len(costs) / 2

	for i := 0; i < len(costs); i++ {
		distance := costs[i][0] - costs[i][1]
		if distance < 0 {
			distance = costs[i][1] - costs[i][0]
		}
		disList[i] = []int{i, distance}
	}

	a := 0
	b := 0
	sort.Sort(DistanceList(disList))
	// fmt.Printf("%v\n", disList)
	for j := 0; j < len(disList); j++ {
		current := disList[j]
		item := costs[current[0]]
		val := item[1]

		if (count > a && item[0] < val) || b >= count {
			val = item[0]
			a++
		} else {
			b++
		}
		// println(current[0], val, count, a)
		target += val
	}

	return target
}
