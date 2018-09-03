package weekly_contest_99

import "testing"

func Test_SurfaceArea(t *testing.T) {
	source := [][]int{
		{2},
	}
	target := surfaceArea(source)
	expect := 10

	if expect != target {
		t.Error("Translate Test_SurfaceArea Fail", expect, target)
	}
	t.Log("Translate Test_SurfaceArea Success")
}
