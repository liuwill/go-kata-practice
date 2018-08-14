package weekly_contest_97

import "testing"

func compareListWithoutOrder(expect []string, target []string) bool {
	if len(expect) != len(target) {
		return false
	}

	count := 0
	for _, exp := range expect {
		for _, tar := range target {
			if exp == tar {
				count++
				break
			}
		}
	}
	return count == len(expect)
}

func Test_UncommonFromSentences(t *testing.T) {
	sourceA := "this apple is sweet"
	sourceB := "this apple is sour"
	target := uncommonFromSentences(sourceA, sourceB)
	expect := []string{"sweet", "sour"}

	if !compareListWithoutOrder(expect, target) {
		t.Error("Translate Test_UncommonFromSentences Fail", target)
	}
	t.Log("Translate Test_UncommonFromSentences Success")
}

func Test_UncommonFromSentencesSecond(t *testing.T) {
	sourceA := "apple apple"
	sourceB := "banana"
	target := uncommonFromSentences(sourceA, sourceB)
	expect := []string{"banana"}

	if !compareListWithoutOrder(expect, target) {
		t.Error("Translate Test_UncommonFromSentencesSecond Fail", target)
	}
	t.Log("Translate Test_UncommonFromSentencesSecond Success")
}
