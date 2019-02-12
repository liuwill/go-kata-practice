package daily_challenge

import "testing"

func compareInterval(intervals []Interval, expect []Interval) bool {
	if len(intervals) != len(expect) {
		return false
	}

	for i, item := range intervals {
		expectItem := expect[i]
		if expectItem.Start != item.Start || expectItem.End != item.End {
			return false
		}
	}
	return true
}

func buildInterval(item []int) Interval {
	return Interval{
		Start: item[0],
		End:   item[1],
	}
}

func buildIntervalList(items [][]int) []Interval {
	target := make([]Interval, len(items))
	for i, val := range items {
		target[i] = buildInterval(val)
	}
	return target
}

func Test_InsertInterval(t *testing.T) {
	source := [][]int{{1, 3}, {6, 9}}
	newItem := []int{2, 5}
	expect := [][]int{{1, 5}, {6, 9}}

	expectInterval := buildIntervalList(expect)
	target := insert(buildIntervalList(source), buildInterval(newItem))
	if !compareInterval(target, expectInterval) {
		t.Error("Translate Test_InsertInterval Fail", expect, target)
	}

	t.Log("Translate Test_InsertInterval Success")
}