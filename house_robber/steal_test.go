package house_robber

import (
	"testing"
)

func Test_Robber(t *testing.T) {
	source := []int{1, 2, 3, 1}
	target := rob(source)
	expect := 4

	if expect != target {
		t.Error("Translate house robber Fail ", target)
	}
	t.Log("Translate Test_Robber Success")
}

func Test_RobberMore(t *testing.T) {
	source := []int{2, 4, 9, 3, 1}
	target := rob(source)
	expect := 12

	if expect != target {
		t.Error("Translate house robber more Fail ", target)
	}
	t.Log("Translate Test_RobberMore Success")
}

func Test_RobberRight(t *testing.T) {
	source := []int{4, 9, 3, 1}
	target := rob(source)
	expect := 10

	if expect != target {
		t.Error("Translate house robber right Fail ", target)
	}
	t.Log("Translate Test_RobberRight Success")
}
