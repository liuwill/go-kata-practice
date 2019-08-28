package weekly_contest

import "testing"

func Test_InvalidTransactions(t *testing.T) {
	sourceCase := [][]string{
		{"alice,20,800,mtv", "alice,50,100,beijing"},
		{"alice,20,800,mtv", "alice,50,1200,mtv"},
		{"alice,20,800,mtv", "bob,50,1200,mtv"},
	}
	expectList := [][]string{
		{"alice,20,800,mtv", "alice,50,100,beijing"},
		{"alice,50,1200,mtv"},
		{"bob,50,1200,mtv"},
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := invalidTransactions(source)
		if !compareStringList(target, expect) {
			t.Error("Run Test_InvalidTransactions Fail", expect, target)
		}
	}

	t.Log("Run Test_InvalidTransactions Success")
}
