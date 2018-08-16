package weekly_contest_94

import (
	"testing"
)

func Test_MinEatingSpeed(t *testing.T) {
	source := []int{3, 6, 7, 11}
	limit := 8
	target := minEatingSpeed(source, limit)
	expect := 4

	if expect != target {
		t.Error("Translate Test_MinEatingSpeed Fail", target)
	}
	t.Log("Translate Test_MinEatingSpeed Success")
}

func Test_MinEatingSpeedSimple(t *testing.T) {
	source := []int{7}
	limit := 1
	target := minEatingSpeed(source, limit)
	expect := 7

	if expect != target {
		t.Error("Translate Test_MinEatingSpeedSimple Fail", target)
	}
	t.Log("Translate Test_MinEatingSpeedSimple Success")
}

func Test_MinEatingSpeedTwo(t *testing.T) {
	source := []int{7, 5}
	limit := 3
	target := minEatingSpeed(source, limit)
	expect := 5

	if expect != target {
		t.Error("Translate Test_MinEatingSpeedTwo Fail", target)
	}
	t.Log("Translate Test_MinEatingSpeedTwo Success")
}

func Test_MinEatingSpeedMore(t *testing.T) {
	source := []int{30, 11, 23, 4, 20}
	limit := 5
	target := minEatingSpeed(source, limit)
	expect := 30

	if expect != target {
		t.Error("Translate Test_MinEatingSpeedMore Fail", target)
	}
	t.Log("Translate Test_MinEatingSpeedMore Success")
}

func Test_MinEatingSpeedMiddle(t *testing.T) {
	source := []int{30, 11, 23, 4, 20}
	limit := 6
	target := minEatingSpeed(source, limit)
	expect := 23

	if expect != target {
		t.Error("Translate Test_MinEatingSpeedError Fail", target)
	}
	t.Log("Translate Test_MinEatingSpeedError Success")
}
