package daily_challenge

import "testing"

func Test_MergeIntervals(t *testing.T) {
	sourceCase := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {2, 3}},
		{{0, 0}},
	}
	expectCase := [][][]int{
		{{1, 6}, {8, 10}, {15, 18}},
		{{1, 5}},
		{{0, 4}},
		{{1, 4}},
		{{0, 0}},
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
