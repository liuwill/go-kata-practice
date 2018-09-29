package weekly_contest_102

import (
	"testing"
)

var (
	subSources = [][]int{
		{3, 1, 2, 4},
	}
	subExpects = []int{17}
)

func Test_SumSubarrayMins(t *testing.T) {
	for i, source := range subSources {
		target := sumSubarrayMins(source)
		expect := subExpects[i]

		if expect != target {
			t.Error("Run Test_SumSubarrayMins Fail", source, target)
		}
	}
	t.Log("Run Test_SumSubarrayMins Success")
}

func Test_SumSubarrayMinsSimple(t *testing.T) {
	for i, source := range subSources {
		target := sumSubarrayMinsSimple(source)
		expect := subExpects[i]

		if expect != target {
			t.Error("Run Test_SumSubarrayMinsSimple Fail", source, target)
		}
	}
	t.Log("Run Test_SumSubarrayMinsSimple Success")
}
