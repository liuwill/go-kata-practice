package stock_trade_fee

import (
	"testing"
)

func Test_maxProfitFee(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	myProfit := maxProfit(prices, fee)
	targetProfit := 8

	if myProfit != targetProfit {
		t.Error("get max profit fee fail")
	}

	t.Log("Test_maxProfitFee pass")
}

func Test_maxProfitFeeCase(t *testing.T) {
	prices := []int{1, 3, 4, 2, 8, 4, 9}
	fee := 2
	myProfit := maxProfit(prices, fee)
	targetProfit := 8

	if myProfit != targetProfit {
		t.Error("get max profit fee case fail")
	}

	t.Log("Test_maxProfitFeeCase pass")
}
