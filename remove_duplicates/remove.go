package remove_duplicates

func removeDuplicates(nums []int) int {
	length := 1
	start := nums[1]
	for i := 1; i < len(nums); i++ {
		if nums[i] != start {
			length++
		}
		start = nums[i]
	}
	return length
}
