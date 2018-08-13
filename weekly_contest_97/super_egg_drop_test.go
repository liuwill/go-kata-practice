package weekly_contest_97

import (
	"testing"
)

func Test_SuperEggDrop(t *testing.T) {
	sourceK := 1
	sourceN := 2
	target := superEggDrop(sourceK, sourceN)
	expect := 2

	if expect != target {
		t.Error("Translate Test_SuperEggDrop Fail", target)
	}
	t.Log("Translate Test_SuperEggDrop Success")
}

func Test_SuperEggDropTwo(t *testing.T) {
	sourceK := 2
	sourceN := 6
	target := superEggDrop(sourceK, sourceN)
	expect := 3

	if expect != target {
		t.Error("Translate Test_SuperEggDropTwo Fail", target)
	}
	t.Log("Translate Test_SuperEggDropTwo Success")
}

func Test_SuperEggDropThree(t *testing.T) {
	sourceK := 3
	sourceN := 14
	target := superEggDrop(sourceK, sourceN)
	expect := 4

	if expect != target {
		t.Error("Translate Test_SuperEggDropThree Fail", target)
	}
	t.Log("Translate Test_SuperEggDropThree Success")
}

func Test_SuperEggDropFail(t *testing.T) {
	sourceK := 1
	sourceN := 3
	target := superEggDrop(sourceK, sourceN)
	expect := 3

	if expect != target {
		t.Error("Translate Test_SuperEggDropFail Fail", target)
	}
	t.Log("Translate Test_SuperEggDropFail Success")
}
