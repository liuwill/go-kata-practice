package weekly_contest_102

import "testing"

func Test_TotalFruit(t *testing.T) {
	source := []int{1, 2, 2, 3}
	target := totalFruit(source)
	expect := 3

	if expect != target {
		t.Error("Run Test_TotalFruit Fail", target)
	}
	t.Log("Run Test_TotalFruit Success")
}

func Test_TotalFruitGroup(t *testing.T) {
	sources := [][]int{
		{1, 2, 1},
		{0, 1, 2, 2},
		{1, 2, 3, 2, 2},
		{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
		{3, 3, 3, 1, 1, 2, 1, 1, 2, 2, 1, 3, 3, 4},
	}
	expects := []int{3, 3, 4, 5, 8}

	for i, source := range sources {
		target := totalFruit(source)
		expect := expects[i]

		if expect != target {
			t.Error("Run Test_TotalFruitGroup Fail", source, target)
		}
	}
	t.Log("Run Test_TotalFruitGroup Success")
}
