package weekly_contest_105

import "testing"

var (
	sources = []string{
		"",
		"==2",
		"ab-cd",
		"a-bC-dEf-ghIj",
		"Test1ng-Leet=code-Q!",
	}
	expects = []string{
		"",
		"==2",
		"dc-ba",
		"j-Ih-gfE-dCba",
		"Qedo1ct-eeLg=ntse-T!",
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
