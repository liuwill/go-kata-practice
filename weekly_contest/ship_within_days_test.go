package weekly_contest

import "testing"

func Test_ShipWithinDays(t *testing.T) {
	sourceCase := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	dayCase := []int{
		5,
	}
	expectList := []int{
		15,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := shipWithinDays(source, dayCase[i])
		if target != expect {
			t.Error("Run Test_ShipWithinDays Fail", expect, target)
		}
	}

	t.Log("Run Test_ShipWithinDays Success")
}
