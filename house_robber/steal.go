package house_robber

func rob(nums []int) int {
	length := len(nums)
	for index := length - 1; index >= 0; index-- {
		left := index + 2
		right := index + 3

		bigger := 0
		if right < length && nums[right] > bigger {
			bigger = nums[right]
		}

		if left < length && nums[left] > bigger {
			bigger = nums[left]
		}
		nums[index] = nums[index] + bigger
	}

	result := 0
	if length >= 1 {
		result = nums[0]
	}

	if length > 1 && nums[1] > nums[0] {
		result = nums[1]
	}

	return result
}
