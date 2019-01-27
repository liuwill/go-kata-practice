package weekly_content

import (
	"testing"
	"time"
)

func compareList(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}

func Test_SortedSquares(t *testing.T) {
	begin := time.Now().Nanosecond()
	source := [][]int{
		{-4, -1, 0, 3, 10},
		{-7, -3, 2, 3, 11},
		{-7, -3, -2, -1, 0},
	}
	expect := [][]int{
		{0, 1, 9, 16, 100},
		{4, 9, 9, 49, 121},
		{0, 1, 4, 9, 49},
	}

	startFirst := time.Now().Nanosecond()
	for i, list := range source {
		short := shortSortedSquares(list)
		if short == nil || !compareList(expect[i], short) {
			t.Error("Translate Test_ShortSortedSquares Fail", short)
		}
	}
	println("Test_ShortSortedSquares", time.Now().Nanosecond()-startFirst)

	start := time.Now().Nanosecond()
	for i, list := range source {
		target := sortedSquares(list)
		if target == nil || !compareList(expect[i], target) {
			t.Error("Translate Test_SortedSquares Fail", target)
		}
	}
	println("Test_SortedSquares", time.Now().Nanosecond()-start)

	startShort := time.Now().Nanosecond()
	for i, list := range source {
		short := shortSortedSquares(list)
		if short == nil || !compareList(expect[i], short) {
			t.Error("Translate Test_ShortSortedSquares Fail", short)
		}
	}
	println("Test_ShortSortedSquares", time.Now().Nanosecond()-startShort)

	t.Log("Translate Test_SortedSquares Success", time.Now().Nanosecond()-begin)
}
