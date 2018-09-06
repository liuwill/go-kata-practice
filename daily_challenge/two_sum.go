package daily_challenge

func twoSum(nums []int, target int) []int {
	result := []int{0, 0}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return result
}
