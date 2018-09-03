package weekly_contest_99

import "testing"

func Test_SurfaceArea(t *testing.T) {
	source := [][]int{
		{1, 2},
		{3, 4},
	}
	target := surfaceArea(source)
	expect := 34

	if expect != target {
		t.Error("Translate Test_SurfaceArea Fail", expect, target)
	}
	t.Log("Translate Test_SurfaceArea Success")
}

func Test_SurfaceAreaZero(t *testing.T) {
	source := [][]int{
		{1, 0},
		{0, 2},
	}
	target := surfaceArea(source)
	expect := 16

	if expect != target {
		t.Error("Translate Test_SurfaceAreaOne Fail", expect, target)
	}
	t.Log("Translate Test_SurfaceAreaOne Success")
}

func Test_SurfaceAreaOne(t *testing.T) {
	source := [][]int{
		{2},
	}
	target := surfaceArea(source)
	expect := 10

	if expect != target {
		t.Error("Translate Test_SurfaceAreaOne Fail", expect, target)
	}
	t.Log("Translate Test_SurfaceAreaOne Success")
}
