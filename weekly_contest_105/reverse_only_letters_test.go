package weekly_contest_105

import "testing"

var (
	sources = []string{
		"ab-cd",
	}
	expects = []string{
		"dc-ba",
	}
)

func Test_ReverseOnlyLetters(t *testing.T) {
	for i, source := range sources {
		target := reverseOnlyLetters(source)
		expect := expects[i]

		if expect != target {
			t.Error("Run Test_ReverseOnlyLettersGroup Fail", source, target)
		}
	}
	t.Log("Run Test_ReverseOnlyLettersGroup Success")
}
