package remove_duplicates

func removeDuplicates(nums []int) int {
	length := 1
	start := nums[0]
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != start {
			length++
			pos = i
		}

		if pos != i {
			nums = append(nums[:i], nums[i+1:]...)

			i--
		}
		start = nums[i]
	}
	return length
}
