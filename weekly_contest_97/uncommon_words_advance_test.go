package weekly_contest_97

import "testing"

func Test_UncommonFromSentencesAdvance(t *testing.T) {
	sourceA := "this apple is sweet"
	sourceB := "this apple is sour"
	target := uncommonFromSentencesAdvance(sourceA, sourceB)
	expect := []string{"sweet", "sour"}

	if !compareListWithoutOrder(expect, target) {
		t.Error("Translate Test_UncommonFromSentencesAdvance Fail", target)
	}
	t.Log("Translate Test_UncommonFromSentencesAdvance Success")
}

func Test_UncommonFromSentencesAdvanceSecond(t *testing.T) {
	sourceA := "apple apple"
	sourceB := "banana"
	target := uncommonFromSentencesAdvance(sourceA, sourceB)
	expect := []string{"banana"}

	if !compareListWithoutOrder(expect, target) {
		t.Error("Translate Test_UncommonFromSentencesAdvanceSecond Fail", target)
	}
	t.Log("Translate Test_UncommonFromSentencesAdvanceSecond Success")
}
