package jump_game

func jump(nums []int) int {
	step := 0
	destination := len(nums)
	index := 0
	for index <= len(nums) {
		size := nums[index]

		nextIndex := 0
		maxLength := 0
		for i := 1; i <= size; i++ {
			position := index + i
			ability := nums[index+i]
			if position >= destination {
				return step + 1
			}

			target := position + ability
			if target > maxLength {
				maxLength = target
				nextIndex = position
			}
		}

		step++
		index = nextIndex
	}

	return step
}
