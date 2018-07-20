package sort_colors

func changePos(nums []int, x int, y int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}

func changBackward(nums []int, start int, pos int) (bool, int) {
	for j := start; j < pos; j++ {
		if nums[j] > nums[pos] {
			changePos(nums, pos, j)
			return true, j
		}
	}
	return false, 0
}

func changForward(nums []int, start int, pos int) (bool, int) {
	for j := start; j > pos; j-- {
		if nums[j] < nums[pos] {
			changePos(nums, pos, j)
			return true, j
		}
	}
	return false, 0
}

func sortColors(nums []int) {
	whitePos := 0
	bluePos := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			isFind, pos := changBackward(nums, whitePos, i)
			if isFind {
				whitePos = pos
				i--
			}
		} else if nums[i] == 2 {
			isFind, pos := changForward(nums, bluePos, i)
			if isFind {
				bluePos = pos
				i--
			}
		}
	}
}
