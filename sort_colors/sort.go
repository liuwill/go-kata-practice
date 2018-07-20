package sort_colors

func changePos(nums []int, x int, y int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}

func sortColors(nums []int) {
	whitePos := 0
	bluePos := len(nums) - 1

	for i, _ := range nums {
		if nums[i] == 0 {
			for j := whitePos; j < i; j++ {
				if nums[j] > nums[i] {
					changePos(nums, i, j)
					whitePos = j
					break
				}
			}
		}

		if nums[i] == 2 {
			for j := bluePos; j > i; j-- {
				if nums[j] < nums[i] {
					changePos(nums, i, j)
					bluePos = j
					break
				}
			}
		}
	}
}
