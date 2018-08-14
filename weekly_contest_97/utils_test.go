package weekly_contest_97

import (
	"testing"
)

func Test_CompareListWithoutOrder(t *testing.T) {
	source1 := []string{"12", "24", "8", "32"}
	source2 := []string{"12", "24"}
	target := compareListWithoutOrder(source1, source2)
	expect := false

	if expect != target {
		t.Error("Translate Test_CompareListWithoutOrder Fail", target)
	}
	t.Log("Translate Test_CompareListWithoutOrder Success")
}
