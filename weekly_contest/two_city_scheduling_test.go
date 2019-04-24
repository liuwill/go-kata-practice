package weekly_contest

import "testing"

func Test_TwoCitySchedCost(t *testing.T) {
	sourceCase := [][][]int{
		[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
		[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}},
	}
	expectList := []int{
		110,
		1859,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := twoCitySchedCost(source)
		if target != expect {
			t.Error("Run Test_TwoCitySchedCost Fail", expect, target)
		}
	}

	t.Log("Run Test_TwoCitySchedCost Success")
}
