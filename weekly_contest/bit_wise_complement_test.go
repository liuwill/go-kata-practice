package weekly_contest

import (
	"testing"
)

func Test_BitwiseComplement(t *testing.T) {
	sourceCase := []int{
		5,
		7,
		10,
		0,
	}
	expectList := []int{
		2,
		0,
		5,
		1,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := bitwiseComplement(source)
		if target != expect {
			t.Error("Run Test_BitwiseComplement Fail", expect, target)
		}
	}

	t.Log("Run Test_BitwiseComplement Success")
}
