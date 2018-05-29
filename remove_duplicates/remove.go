package remove_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	length := 1
	capacity := len(nums)
	start := nums[0]
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != start {
			length++
			pos = i
		}

		if pos != i {
			nums = append(nums[:i], nums[i+1:]...)
			capacity--
			i--
		}
		start = nums[i]
	}

	return length
}
