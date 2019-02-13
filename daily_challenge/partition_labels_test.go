package daily_challenge

import "testing"

func Test_PartitionLabels(t *testing.T) {
	sourceCase := []string{
		"ababcbacadefegdehijhklij",
	}
	expectCase := [][]int{
		{9, 7, 8},
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
