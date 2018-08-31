package daily_challenge

const (
	Radiant = "Radiant"
	Dire    = "Dire"
)

func predictPartyVictory(senate string) string {
	voteCount := map[string]int{
		"R": 0,
		"D": 0,
	}
	otherMap := map[string]string{
		"R": "D",
		"D": "R",
	}

	left := senate
	victory := ""
	for len(victory) == 0 {
		current := ""
		leftCount := map[string]int{
			"R": 0,
			"D": 0,
		}
		for _, v := range left {
			side := string(v)
			if voteCount[side] > 0 {
				voteCount[side]--
			} else {
				current += side
				leftCount[side]++
				other := otherMap[side]
				voteCount[other]++
			}
		}

		if leftCount["R"] > 0 && leftCount["D"] > 0 {
			left = current
			continue
		}

		if leftCount["R"] == 0 {
			victory = Dire
		} else {
			victory = Radiant
		}
	}

	return victory
}
