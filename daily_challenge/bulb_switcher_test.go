package daily_challenge

import (
	"testing"
)

func Test_BulbSwitch(t *testing.T) {
	sourceCase := []int{
		3,
		5,
	}
	expectList := []int{
		1,
		2,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := bulbSwitch(source)
		if target != expect {
			t.Error("Run Test_BulbSwitch Fail", source, expect, target)
		}
	}

	t.Log("Run Test_BulbSwitch Success")
}
