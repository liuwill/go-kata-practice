package daily_challenge

/**
 * daily-challenge-268
 * PUZZLE: Missing Number
 */
func missingNumber(nums []int) int {
	target := 0
	expect := 0
	for i, val := range nums {
		target += i + 1
		expect += val
	}
	return target - expect
}
