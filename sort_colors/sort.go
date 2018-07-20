package sort_colors

func sortColors(nums []int) {
	for i, _ := range nums {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
}
