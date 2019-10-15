package weekly_contest

import (
	"testing"
)

func comparePointMatch(source [][]int, target [][]int) bool {
	if len(source) != len(target) {
		return false
	}
	targetCount := 0
	targetPos := make([]int, len(source))

	for _, item := range source {
		for pos, matcher := range target {
			if item[0] == matcher[0] && item[1] == matcher[1] {
				targetCount++
				targetPos[pos]++
			}
		}
	}

	if targetCount != len(source) {
		return false
	}

	for _, t := range targetPos {
		if t != 1 {
			return false
		}
	}

	return true
}

func Test_QueensAttacktheKing(t *testing.T) {
	queueCase := [][][]int{
		{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
		{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}},
		{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}},
	}
	kingCase := [][]int{
		{0, 0},
		{3, 3},
		{3, 4},
	}
	expectList := [][][]int{
		{{0, 1}, {1, 0}, {3, 3}},
		{{2, 2}, {3, 4}, {4, 4}},
		{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}},
	}

	for i, queues := range queueCase {
		expect := expectList[i]

		target := queensAttacktheKing(queues, kingCase[i])
		if !comparePointMatch(target, expect) {
			t.Error("Run Test_QueensAttacktheKing Fail", expect, target)
		}
	}

	t.Log("Run Test_QueensAttacktheKing Success")
}
