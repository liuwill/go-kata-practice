package daily_challenge

import "testing"

func Test_PartitionLabels(t *testing.T) {
	sourceCase := []string{
		"ababcbacadefegdehijhklij",
		"vvvvooopopttttk",
		"vvvvooopopttttkk",
		"p",
		"ttt",
		"tkl",
		"tklk",
		"tktlk",
		"eccbbbbdec",
	}
	expectCase := [][]int{
		{9, 7, 8},
		{4, 6, 4, 1},
		{4, 6, 4, 2},
		{1},
		{3},
		{1, 1, 1},
		{1, 3},
		{5},
		{10},
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		target := partitionLabels(source)
		if !compareList(target, expect) {
			t.Error("Translate Test_PartitionLabels Fail", expect, target)
		}
	}

	t.Log("Translate Test_PartitionLabels Success")
}
