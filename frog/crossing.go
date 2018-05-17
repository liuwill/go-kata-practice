package frog

type JumpAction struct {
	position int
	action   int
}

func getPossibleDistance(distance int, position int) []int {
	var target []int
	if position == 1 {
		target = []int{1}
	} else {
		target = []int{}
		if distance > 1 {
			target = append(target, distance-1)
		}
		target = append(target, distance, distance+1)
	}
	return target
}

func findFront(stones []int, index int, distance int) int {
	i := index - 1
	front := stones[index] - distance
	for i >= 0 {
		if stones[i] == front {
			if i > 1 || (i <= 1 && distance <= i+1) {
				return i
			}
		} else if stones[i] < front {
			break
		}

		i--
	}
	return -1
}

func canCross(stones []int) bool {
	if len(stones) < 2 {
		return true
	}

	// distance := 0
	stack := []JumpAction{}
	lastPosition := len(stones) - 1
	index := lastPosition - 1
	stack = []JumpAction{
		JumpAction{
			position: lastPosition,
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
		// println("+++", currentIndex, stones[currentIndex], currentAction, len(stack), len(possibleDistances))
		for _, action := range possibleDistances {
			frontPosition := findFront(stones, currentIndex, action)

			if frontPosition >= 0 {
				// println(len(stack), "====", currentIndex, currentJump.distance, currentAction, ":", frontPosition, stones[frontPosition], action)
				stack = append(stack, JumpAction{
					position: frontPosition,
					action:   action,
				})
			}
		}
	}
	return false
}
