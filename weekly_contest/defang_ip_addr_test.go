package weekly_contest

import "testing"

func Test_DefangIPaddr(t *testing.T) {
	sourceCase := []string{
		"1.1.1.1",
		"255.100.50.0",
	}
	expectList := []string{
		"1[.]1[.]1[.]1",
		"255[.]100[.]50[.]0",
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := defangIPaddr(source)
		if target != expect {
			t.Error("Run Test_DefangIPaddr Fail", expect, target)
		}
	}

	t.Log("Run Test_DefangIPaddr Success")
}
