package frog

func getPossibleDistance(distance int, position int) []int {
	target := []int{}
	if position == 1 {
		target = append(target, 1)
	} else {
		target = append(target, distance+1, distance)
		if distance > 1 {
			target = append(target, distance-1)
		}
	}
	return target
}

type JumpAction struct {
	position int
	distance int
	action   []int
}

func couldCross(stones *[]int, index int, front int) bool {

}

func canCross(stones []int) bool {
	if len(stones) < 3 {
		return true
	}

	// distance := 0
	stack := []JumpAction{}
	lastPosition := len(stones) - 1
	index := lastPosition - 1
	stack = append(stack, JumpAction{
		position: lastPosition,
		distance: stones[lastPosition],
		action:   []int{stones[lastPosition] - stones[index]},
	})

	for len(stack) > 0 {
		currentJump := stack[0]
		stack = stack[1:]

		currentAction := currentJump.action
		currentIndex := currentJump.position
		if currentIndex == 0 {
			return true
		}

		for _, action := range currentAction {
			if couldCross(&stones, currentIndex, action) {

			}
		}
	}
	return false
}
