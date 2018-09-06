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

func twoSumWithMap(nums []int, target int) []int {
	result := []int{0, 0}

	posMap := make(map[int]int)
	posMap[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		left := target - nums[i]
		if _, ok := posMap[left]; ok {
			result = []int{posMap[left], i}
			break
		}
		posMap[nums[i]] = i
	}

	return result
}
