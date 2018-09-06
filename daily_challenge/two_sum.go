package daily_challenge

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)

	find := false
	for i := 0; i < len(nums)-1 && !find; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] != target {
				continue
			}

			result[0] = nums[i]
			result[1] = nums[j]
			find = true
		}
	}

	return result
}
