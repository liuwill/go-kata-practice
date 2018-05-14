package jump_game

func jump(nums []int) int {
	step := 0
	if len(nums) <= 1 {
		return step
	} else if len(nums)-1 <= nums[0] {
		return step + 1
	}

	destination := len(nums) - 1
	index := 0
	for index < destination {
		size := nums[index]

		nextIndex := 0
		maxLength := 0
		for i := 1; i <= size; i++ {
			position := index + i
			ability := nums[position]

			target := position + ability
			if target > maxLength {
				maxLength = target
				nextIndex = position
			}

			if target >= destination {
				break
			}
		}

		step++
		index = nextIndex

		if maxLength >= destination {
			step++
			index = destination
		}
	}

	return step
}
