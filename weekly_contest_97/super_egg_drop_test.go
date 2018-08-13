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
		t.Error("Translate Test_SuperEggDrop Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDrop Success")
}

func Test_SuperEggDropTwo(t *testing.T) {
	sourceK := 2
	sourceN := 6
	target := superEggDrop(sourceK, sourceN)
	expect := 3

	if expect != target {
		t.Error("Translate Test_SuperEggDropTwo Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropTwo Success")
}

func Test_SuperEggDropThree(t *testing.T) {
	sourceK := 3
	sourceN := 14
	target := superEggDrop(sourceK, sourceN)
	expect := 4

	if expect != target {
		t.Error("Translate Test_SuperEggDropThree Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropThree Success")
}

func Test_SuperEggDropFail(t *testing.T) {
	sourceK := 1
	sourceN := 3
	target := superEggDrop(sourceK, sourceN)
	expect := 3

	if expect != target {
		t.Error("Translate Test_SuperEggDropFail Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropFail Success")
}

func Test_SuperEggDropError(t *testing.T) {
	sourceK := 2
	sourceN := 7
	target := superEggDrop(sourceK, sourceN)
	expect := 4

	if expect != target {
		t.Error("Translate Test_SuperEggDropFail Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropFail Success")
}

func Test_SuperEggDropEasy(t *testing.T) {
	sourceK := 2
	sourceN := 1
	target := superEggDrop(sourceK, sourceN)
	expect := 1

	if expect != target {
		t.Error("Translate Test_SuperEggDropEasy Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropEasy Success")
}

func Test_SuperEggDropMore(t *testing.T) {
	sourceK := 2
	sourceN := 3
	target := superEggDrop(sourceK, sourceN)
	expect := 2

	if expect != target {
		t.Error("Translate Test_SuperEggDropMore Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropMore Success")
}

func Test_SuperEggDropEqual(t *testing.T) {
	sourceK := 2
	sourceN := 2
	target := superEggDrop(sourceK, sourceN)
	expect := 2

	if expect != target {
		t.Error("Translate Test_SuperEggDropEqual Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropEqual Success")
}

func Test_SuperEggDropFour(t *testing.T) {
	sourceK := 2
	sourceN := 4
	target := superEggDrop(sourceK, sourceN)
	expect := 3

	if expect != target {
		t.Error("Translate Test_SuperEggDropFour Fail", expect, target)
	}
	t.Log("Translate Test_SuperEggDropFour Success")
}
