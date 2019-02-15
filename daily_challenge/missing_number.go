package daily_challenge

func missingNumber(nums []int) int {
	target := 0
	expect := 0
	for i, val := range nums {
		target += i + 1
		expect += val
	}
	return target - expect
}
