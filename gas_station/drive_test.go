package gas_station

import (
	"testing"
)

func Test_CanDriveCompleteSuccess(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	target := canDriveCompleteCircuit(gas, cost)
	expect := 3

	if expect != target {
		t.Error("Translate Gas Station CanDriveSuccess Fail ", target)
	}
	t.Log("Translate Test_DriveCanCompleteSuccess Success")
}

func Test_CanDriveCompleteFail(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	target := canDriveCompleteCircuit(gas, cost)
	expect := -1

	if expect != target {
		t.Error("Translate Gas Station CanDriveFail Fail ", target)
	}
	t.Log("Translate Test_CanDriveCompleteFail more Success")
}
