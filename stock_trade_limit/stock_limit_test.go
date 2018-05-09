package stock_trade_limit

import (
	"testing"
)

func Test_maxProfitLimit(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	myProfit := maxProfit(prices)
	targetProfit := 6
	if myProfit != targetProfit {
		t.Error("get max profit limit fail")
	}

	t.Log("Test_maxProfitLimit pass")
}
