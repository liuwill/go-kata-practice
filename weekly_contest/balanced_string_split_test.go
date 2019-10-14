package weekly_contest

import "testing"

func Test_BalancedStringSplit(t *testing.T) {
	sourceCase := []string{
		"RLRRLLRLRL", "RLLLLRRRLR", "LLLLRRRR",
	}
	expectList := []int{
		4, 3, 1,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := balancedStringSplit(source)
		if target != expect {
			t.Error("Run Test_BalancedStringSplit Fail", expect, target)
		}
	}

	t.Log("Run Test_BalancedStringSplit Success")
}
