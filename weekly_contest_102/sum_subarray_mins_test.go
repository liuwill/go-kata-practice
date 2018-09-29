package weekly_contest_102

import (
	"math/rand"
	"testing"
	"time"
)

var (
	subSources = [][]int{
		{3, 1, 2, 4},
		{3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4, 3, 1, 2, 4},
	}
	subExpects = []int{17, 1733}
)

func generateRandomArray(total int) []int {
	list := []int{}
	for i := 0; i < total; i++ {
		list = append(list, rand.Intn(MAX_SUBARRAY_LENGTH-1)+1)
	}
	return list
}

func Test_SumSubarrayMinsBanchmark(t *testing.T) {
	list := generateRandomArray(500)
	println(len(list))

	startSimple := time.Now().UnixNano()
	target := sumSubarrayMinsFast(list)
	endSimple := time.Now().UnixNano()
	println("sumSubarrayMinsFast", startSimple, endSimple, endSimple-startSimple)

	start := time.Now().UnixNano()
	expect := sumSubarrayMins(list)
	end := time.Now().UnixNano()
	println("sumSubarrayMins", start, end, end-start)

	if expect != target {
		t.Error("Run Test_SumSubarrayMinsBanchmark Fail", expect, target)
	}
	t.Log("Run Test_SumSubarrayMinsBanchmark Success")
}

func Test_SumSubarrayMins(t *testing.T) {
	for i, source := range subSources {
		target := sumSubarrayMins(source)
		expect := subExpects[i]

		if expect != target {
			t.Error("Run Test_SumSubarrayMins Fail", i, expect, target)
		}
	}
	t.Log("Run Test_SumSubarrayMins Success")
}

func Test_SumSubarrayMinsFast(t *testing.T) {
	for i, source := range subSources {
		target := sumSubarrayMinsFast(source)
		expect := subExpects[i]

		if expect != target {
			t.Error("Run Test_SumSubarrayMinsFast Fail", i, expect, target)
		}
	}
	t.Log("Run Test_SumSubarrayMinsFast Success")
}
