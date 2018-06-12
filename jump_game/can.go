package jump_game

func canJump(nums []int) bool {
	target := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		current := nums[i]
		if i+current >= target {
			target = i
		}
	}

	return target == 0
}
