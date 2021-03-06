package weekly_contest

import "testing"

func Test_NumRookCaptures(t *testing.T) {
	sourceCase := [][][]byte{
		{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'R', '.', '.', '.', 'p'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
		{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'B', 'R', 'B', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
		{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'p', 'p', '.', 'R', '.', 'p', 'B', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
		{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', 'B', 'B', 'B', 'B', 'B', '.'}, {'.', 'p', 'B', 'p', 'p', 'p', 'B', 'p'}, {'.', 'p', 'B', 'p', 'R', 'p', 'B', 'p'}, {'.', 'p', 'B', 'p', 'p', 'p', 'B', 'p'}, {'.', '.', 'B', 'B', 'B', 'B', 'B', '.'}, {'.', '.', '.', 'p', 'p', 'p', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}},
	}
	expectList := []int{
		3,
		0,
		3,
		4,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := numRookCaptures(source)
		if target != expect {
			t.Error("Run Test_NumRookCaptures Fail", expect, target)
		}
	}

	t.Log("Run Test_NumRookCaptures Success")
}
