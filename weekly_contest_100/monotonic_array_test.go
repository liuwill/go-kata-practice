package weekly_contest_100

import "testing"

func Test_IsMonotonic(t *testing.T) {
	source := []int{1, 2, 2, 3}
	target := isMonotonic(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsMonotonic Fail", target)
	}
	t.Log("Translate Test_IsMonotonic Success")
}

func Test_IsMonotonicSmall(t *testing.T) {
	source := []int{1}
	target := isMonotonic(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsMonotonicSmall Fail", target)
	}
	t.Log("Translate Test_IsMonotonicSmall Success")
}

func Test_IsMonotonicDecrease(t *testing.T) {
	source := []int{6, 5, 4, 4}
	target := isMonotonic(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsMonotonicDecrease Fail", target)
	}
	t.Log("Translate Test_IsMonotonicDecrease Success")
}

func Test_IsMonotonicFail(t *testing.T) {
	source := []int{1, 3, 2}
	target := isMonotonic(source)
	expect := false

	if expect != target {
		t.Error("Translate Test_IsMonotonicFail Fail", target)
	}
	t.Log("Translate Test_IsMonotonicFail Success")
}
