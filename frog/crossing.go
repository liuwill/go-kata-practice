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
	action   int
}

func findFront(stones *[]int, index int, front int) int {
	return -1
}

func canCross(stones []int) bool {
	if len(stones) < 3 {
		return true
	}

	// distance := 0
	stack := []JumpAction{}
	lastPosition := len(stones) - 1
	index := lastPosition - 1
	stack = []JumpAction{
		JumpAction{
			position: lastPosition,
			distance: stones[lastPosition],
			action:   stones[lastPosition] - stones[index],
		},
	}

	for len(stack) > 0 {
		lastPos := len(stack) - 1
		currentJump := stack[lastPos]
		stack = stack[:lastPos]

		currentAction := currentJump.action
		currentIndex := currentJump.position
		if currentIndex == 0 {
			return true
		}

		possibleDistances := getPossibleDistance(currentAction, currentIndex)
		for _, action := range possibleDistances {
			frontPosition := findFront(&stones, currentIndex, action)
			if frontPosition == 0 {
				return true
			}

			if frontPosition >= 0 {
				stack = append(stack, JumpAction{
					position: currentIndex,
					distance: stones[currentIndex],
					action:   stones[currentIndex] - stones[frontPosition],
				})
			}
		}
	}
	return false
}
