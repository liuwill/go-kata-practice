package stock_trade_limit

import (
	"testing"
)

func Test_maxProfitLimit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	myProfit := maxProfit(prices)
	targetProfit := 5
	if myProfit != targetProfit {
		t.Error("get max profit limit fail")
	}

	t.Log("Test_maxProfitLimit pass")
}

func Test_maxProfitLimitCase(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	myProfit := maxProfit(prices)
	targetProfit := 0
	if myProfit != targetProfit {
		t.Error("get max profit limit case fail")
	}

	t.Log("Test_maxProfitLimitCase pass")
}

func Test_maxProfitFeeEmpty(t *testing.T) {
	prices := []int{1}
	myProfit := maxProfit(prices)
	targetProfit := 0

	if myProfit != targetProfit {
		t.Error("get max profit empty fail")
	}

	t.Log("Test_maxProfitFeeEmpty pass")
}
