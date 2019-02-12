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

func Test_MergeIntervals(t *testing.T) {
	sourceCase := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
	}
	expectCase := [][][]int{
		{{1, 6}, {8, 10}, {15, 18}},
		{{1, 5}},
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		expectInterval := buildIntervalList(expect)
		target := merge(buildIntervalList(source))
		if !compareInterval(target, expectInterval) {
			t.Error("Translate Test_MergeIntervals Fail", expect, target)
		}
	}

	t.Log("Translate Test_MergeIntervals Success")
}
