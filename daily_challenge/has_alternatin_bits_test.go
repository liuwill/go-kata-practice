package daily_challenge

import (
	"testing"
)

func Test_HasAlternatingBits(t *testing.T) {
	sourceCase := []int{
		1,
		5,
		7,
		11,
		10,
	}
	expectList := []bool{
		true,
		true,
		false,
		false,
		true,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := hasAlternatingBits(source)
		if target != expect {
			t.Error("Run Test_HasAlternatingBits Fail", source, expect, target)
		}
	}

	t.Log("Run Test_HasAlternatingBits Success")
}
