package gas_station

import (
	"testing"
)

func Test_CanCompleteSuccess(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	target := canCompleteCircuit(gas, cost)
	expect := 3

	if expect != target {
		t.Error("Translate Gas Station CanSuccess Fail ", target)
	}
	t.Log("Translate Test_CanCompleteSuccess Success")
}

func Test_CanCompleteFail(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	target := canCompleteCircuit(gas, cost)
	expect := -1

	if expect != target {
		t.Error("Translate Gas Station CanFail Fail ", target)
	}
	t.Log("Translate Test_CanCompleteFail more Success")
}
