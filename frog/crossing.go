package frog

func getPossibleDistance(distance int, position int) []int {
	target := []int{}
	if position == 1 {
		target = append(target, 1)
	} else {
		if distance > 1 {
			target = append(target, distance-1)
		}
		target = append(target, distance, distance+1)
	}
	return target
}

type JumpAction struct {
	position int
	distance int
	action   int
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
		} else if currentIndex == 1 && currentAction > 2 {
			continue
		}

		possibleDistances := getPossibleDistance(currentAction, currentIndex)
		println("+++", currentIndex, currentJump.distance, currentAction, len(stack), len(possibleDistances))
		for _, action := range possibleDistances {
			frontPosition := findFront(stones, currentIndex, action)

			if frontPosition >= 0 {
				// println(len(stack), "====", currentIndex, currentJump.distance, currentAction, ":", frontPosition, stones[frontPosition], action)
				stack = append(stack, JumpAction{
					position: frontPosition,
					distance: stones[frontPosition],
					action:   action,
				})
			}
		}
	}
	return false
}
