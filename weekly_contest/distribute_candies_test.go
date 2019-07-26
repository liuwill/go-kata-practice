package weekly_contest

import "testing"

func Test_DistributeCandies(t *testing.T) {
	candyCase := []int{
		7, 10,
	}
	peopleCase := []int{
		4,
		3,
	}
	expectList := [][]int{
		[]int{1, 2, 3, 1},
		[]int{5, 2, 3},
	}

	for i, candy := range candyCase {
		people := peopleCase[i]
		expect := expectList[i]

		target := distributeCandies(candy, people)
		if !compareList(target, expect) {
			t.Error("Run Test_DistributeCandies Fail", expect, target)
		}
	}

	t.Log("Run Test_DistributeCandies Success")
}
