package weekly_contest

import "testing"

func Test_CountVowelPermutation(t *testing.T) {
	sourceCase := []int{
		1, 2, 5,
		// 144,
	}
	expectList := []int{
		5, 10, 68,
		// 18208803,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := countVowelPermutation(source)
		if target != expect {
			t.Error("Run Test_CountVowelPermutation Fail", expect, target)
		}
	}

	t.Log("Run Test_CountVowelPermutation Success")
}
