package weekly_contest_93

import (
	"testing"
)

func Test_MinRefuelStops(t *testing.T) {
	destination := 100
	startFuel := 10
	stations := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	target := minRefuelStops(destination, startFuel, stations)
	expect := 2

	if expect != target {
		t.Error("Translate minRefuelStops Fail", target)
	}
	t.Log("Translate Test_MinRefuelStops Success")
}
