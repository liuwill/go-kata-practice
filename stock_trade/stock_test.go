package stock_trade

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	myProfit := maxProfit(prices)
	targetProfit := 7
	if myProfit != targetProfit {
		t.Error("get max profit fail")
	}

	t.Log("Test_maxProfit pass")
}

func Test_maxProfitLine(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	myProfit := maxProfit(prices)
	targetProfit := 4
	if myProfit != targetProfit {
		t.Error("get max profit line fail")
	}

	t.Log("Test_maxProfitLine pass")
}

func Test_maxProfitZero(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	myProfit := maxProfit(prices)
	targetProfit := 0
	if myProfit != targetProfit {
		t.Error("get max profit zero fail")
	}

	t.Log("Test_maxProfitZero pass")
}
