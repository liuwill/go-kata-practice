package weekly_contest

import "testing"

func Test_DistanceBetweenBusStops(t *testing.T) {
	distanceCase := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	pointList := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
	}
	expectList := []int{1, 3, 4}

	for i, distanceList := range distanceCase {
		expect := expectList[i]
		point := pointList[i]

		target := distanceBetweenBusStops(distanceList, point[0], point[1])
		if target != expect {
			t.Error("Run Test_DistanceBetweenBusStops Fail", expect, target)
		}
	}

	t.Log("Run Test_DistanceBetweenBusStops Success")
}
