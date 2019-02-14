package daily_challenge

/**
 * daily-challenge-287
 * PUZZLE: Find the Duplicate Number
 */
func findDuplicate(nums []int) int {
	target := 0
	for i, val := range nums {
		isFind := false
		j := i + 1
		for ; j < len(nums); j++ {
			if val == nums[j] {
				isFind = true
				break
			}
		}

		if isFind {
			target = val
			break
		}
	}
	return target
}
