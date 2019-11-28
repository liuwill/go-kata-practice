package weekly_contest

import (
	"testing"
)

func compareMatrixArray(target [][]string, expect [][]string) bool {
	if len(target) != len(expect) {
		return false
	}

	for i, _ := range target {
		if len(target[i]) != len(expect[i]) {
			return false
		}

		for j, _ := range target[i] {
			if target[i][j] != expect[i][j] {
				return false
			}
		}
	}

	return true
}

func Test_SuggestedProducts(t *testing.T) {
	productCase := [][]string{
		{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
		{"havana"},
		{"bags", "baggage", "banner", "box", "cloths"},
		{"havana"},
	}
	wordCase := []string{"mouse", "havana", "bags", "tatiana"}
	expectList := [][][]string{
		{
			{"mobile", "moneypot", "monitor"},
			{"mobile", "moneypot", "monitor"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
		},
		{
			{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"},
		},
		{
			{"baggage", "bags", "banner"},
			{"baggage", "bags", "banner"},
			{"baggage", "bags"},
			{"bags"},
		},
		{
			{}, {}, {}, {}, {}, {}, {},
		},
	}

	for i, products := range productCase {
		expect := expectList[i]

		target := suggestedProducts(products, wordCase[i])
		if !compareMatrixArray(target, expect) {
			t.Error("Run Test_SuggestedProducts Fail", expect, target)
		}
	}

	t.Log("Run Test_SuggestedProducts Success")
}
