package weekly_contest

import "testing"

func Test_MaxScoreSightseeingPair(t *testing.T) {
	sourceCase := [][]int{
		{8, 1, 5, 2, 6},
	}
	expectList := []int{
		11,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := maxScoreSightseeingPair(source)
		if target != expect {
			t.Error("Run Test_MaxScoreSightseeingPair Fail", expect, target)
		}

		targetFast := maxScoreSightseeingPairFast(source)
		if target != expect {
			t.Error("Run Test_MaxScoreSightseeingPair maxScoreSightseeingPairFast Fail", expect, targetFast)
		}
	}

	t.Log("Run Test_MaxScoreSightseeingPair Success")
}
