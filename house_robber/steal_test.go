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
