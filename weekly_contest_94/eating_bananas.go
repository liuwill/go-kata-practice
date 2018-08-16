package weekly_contest_94

func countEatableHour(K int, piles []int) int {
	hours := 0
	for _, v := range piles {
		hours += v / K
		if v%K > 0 {
			hours++
		}
	}

	return hours
}

func minEatingSpeed(piles []int, H int) int {
	largest := piles[0]
	smallest := 1
	answer := 0
	if len(piles) == 1 {
		answer = piles[0] / H
		if piles[0]%H > 0 {
			answer++
		}
		return answer
	}

	total := 0
	for _, v := range piles {
		if v > largest {
			largest = v
		}
		total += v
	}
	smallest = total / H

	middle := (largest + smallest) / 2
	top := smallest
	bottom := largest
	for middle >= smallest && middle <= largest && top != bottom {
		current := countEatableHour(middle, piles)
		if current > H {
			top = middle
		} else {
			bottom = middle
			if middle < answer || answer == 0 {
				answer = middle
			} else if middle >= answer {
				break
			}
		}

		center := (top + bottom) / 2
		if center != middle {
			middle = center
		} else {
			middle = center + 1
		}

		// println(top, bottom, center, middle, answer, current, largest, smallest)
	}
	return answer
}
