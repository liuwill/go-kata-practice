package stock_trade

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	myProfit := maxProfit(prices)
	targetProfit := 7
	if myProfit == targetProfit {
		t.Error("get max profit")
	}

	t.Log("Test_maxProfit pass")
}
