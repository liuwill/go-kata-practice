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
		for _, v := range left {
			side := string(v)
			if voteCount[side] > 0 {
				voteCount[side]--
			} else {
				current += side
				other := otherMap[side]
				voteCount[other]++
			}
		}

		if voteCount["R"] < voteCount["D"] {
			victory = Radiant
		} else if voteCount["R"] > voteCount["D"] {
			victory = Dire
		}
		left = current
	}

	return victory
}
