package rain

func buildGoNextStep(direction bool) func(int) int {
	if direction {
		return func(x int) int {
			return x + 1
		}
	}
	return func(x int) int {
		return x - 1
	}
}

func findMax(height []int, start int, end int, direction bool) int {
	max := 0
	position := -1

	index := start
	goStep := buildGoNextStep(direction)
	hasMore := func(index int) bool {
		if direction {
			return index <= end
		}
		return index >= start
	}
	// goStep := func(x int) int { return x + 1 }
	// hasMore := func(index int) bool { return index <= end }
	if !direction {
		index = end
		// goStep = func(x int) int { return x - 1 }
		// hasMore = func(index int) bool { return index >= start }
	}

	for {
		if !hasMore(index) {
			break
		}

		current := height[index]
		if current > max {
			max = current
			position = index
		}
		index = goStep(index)
	}

	return position
}

func trapWater(height []int, start int, end int, small int) int {
	water := 0
	if end-start < 2 {
		return water
	}

	for i := start + 1; i < end; i++ {
		water = water + (small - height[i])
	}
	return water
}

func isBorder(height []int, current int, direction bool) bool {
	if direction && current >= len(height)-1 {
		return true
	} else if !direction && current <= 0 {
		return true
	}

	return false
}

func countBlockWater(height []int, previousHeighest int, direction bool) int {
	water := 0

	for {
		if isBorder(height, previousHeighest, direction) {
			break
		}

		frontHeighest := -1
		if !direction {
			frontHeighest = findMax(height, 0, previousHeighest-1, direction)
		} else {
			frontHeighest = findMax(height, previousHeighest+1, len(height)-1, direction)
		}

		if frontHeighest < 0 {
			break
		}

		unitWater := 0
		if !direction {
			unitWater = trapWater(height, frontHeighest, previousHeighest, height[frontHeighest])
		} else {
			unitWater = trapWater(height, previousHeighest, frontHeighest, height[frontHeighest])
		}
		water = water + unitWater
		previousHeighest = frontHeighest
	}

	return water
}

func trap(height []int) int {
	water := 0
	if len(height) < 3 {
		return water
	}
	heighest := findMax(height, 0, len(height)-1, true)

	water = water + countBlockWater(height, heighest, false)
	water = water + countBlockWater(height, heighest, true)
	/*
			previousHeighest := heighest
			for {
				if previousHeighest <= 0 {
					break
				}
				frontHeighest := findMax(height, 0, previousHeighest-1, false)
				if frontHeighest < 0 {
					break
				}

				water = water + trapWater(height, frontHeighest, previousHeighest, height[frontHeighest])
				previousHeighest = frontHeighest
			}

			previousHeighest = heighest
			for {
				if previousHeighest >= len(height)-1 {
					break
				}
				frontHeighest := findMax(height, previousHeighest+1, len(height)-1, true)
				if frontHeighest < 0 {
					break
				}

				water = water + trapWater(height, previousHeighest, frontHeighest, height[frontHeighest])
				previousHeighest = frontHeighest
		  }
	*/

	return water
}
