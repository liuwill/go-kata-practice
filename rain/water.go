package rain

func findMax(height []int, start int, end int, direction bool) int {
	max := 0
	position := -1

	index := start
	goStep := func(x int) int { return x + 1 }
	hasMore := func(index int) bool { return index <= end }
	if !direction {
		index = end
		goStep = func(x int) int { return x - 1 }
		hasMore = func(index int) bool { return index >= start }
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

/*
func countBlockWater(height []int, previousHeighest int, direction bool) int {
	water := 0
	for {
		if previousHeighest <= 0 {
			break
		}
		frontHeighest := findMax(height, 0, previousHeighest-1, direction)
		if frontHeighest < 0 {
			break
		}

		water = water + trapWater(height, frontHeighest, previousHeighest, height[frontHeighest])
		previousHeighest = frontHeighest
	}

	return water
}
*/

func trap(height []int) int {
	water := 0
	if len(height) < 3 {
		return water
	}
	heighest := findMax(height, 0, len(height)-1, true)

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

	return water
}
