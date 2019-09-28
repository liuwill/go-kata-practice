package weekly_contest

import "sort"

const MAX = 1000000

/**
 * daily-challenge-1200
 * PUZZLE: Minimum Absolute Difference
 */
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	subList := make([]int, len(arr)-1)
	min := MAX
	for i := 0; i < len(arr)-1; i++ {
		subList[i] = arr[i+1] - arr[i]
		if subList[i] < min {
			min = subList[i]
		}
	}

	target := [][]int{}
	for i, val := range subList {
		if val == min {
			target = append(target, []int{arr[i], arr[i+1]})
		}
	}

	return target
}
