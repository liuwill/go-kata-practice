package daily_challenge

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)

	find := false
	for i := 0; i < len(nums)-1 && !find; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] != target {
				continue
			}

			result[0] = i
			result[1] = j
			find = true
		}
	}

	return result
}
