package keys_rooms

import (
	"testing"
)

func Test_Visit(t *testing.T) {
	source := [][]int{{1}, {2}, {3}, {}}
	target := canVisitAllRooms(source)
	expect := true

	if expect != target {
		t.Error("Translate visit Fail", target)
	}
	t.Log("Translate Test_Visit Success")
}

func Test_VisitFail(t *testing.T) {
	source := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	target := canVisitAllRooms(source)
	expect := false

	if expect != target {
		t.Error("Translate visit false Fail", target)
	}
	t.Log("Translate Test_VisitFail Success")
}
