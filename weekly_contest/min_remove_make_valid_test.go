package weekly_contest

import "testing"

func Test_MinRemoveToMakeValid(t *testing.T) {
	sourceCase := []string{
		"lee(t(c)o)de)",
		"a)b(c)d",
		"))((",
		"(a(b(c)d)",
	}
	expectList := []string{
		"lee(t(c)o)de",
		"ab(c)d",
		"",
		"a(b(c)d)",
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := minRemoveToMakeValid(source)
		if target != expect {
			t.Error("Run Test_MinRemoveToMakeValid Fail", i, expect, target)
		}
	}

	t.Log("Run Test_MinRemoveToMakeValid Success")
}
