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

	for _, v := range senate {
		side := string(v)
		if voteCount[side] > 0 {
			voteCount[side]--
		} else {
			other := otherMap[side]
			voteCount[other]++
		}
	}

	if voteCount["R"] < voteCount["D"] {
		return Radiant
	}
	return Dire
}
