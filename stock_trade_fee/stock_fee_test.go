package stock_trade_fee

import (
	"testing"
)

func Test_FeeAccount(t *testing.T) {
	myAccount := Account{
		fee:    2,
		profit: 0,
		cost:   0,
		hold:   false,
	}

	myAccount.detectCost(3)
	if myAccount.cost != 3 {
		t.Error("test fee account fail")
	}
	myAccount.detectCost(4)
	if myAccount.cost == 4 {
		t.Error("test fee account fail")
	}
	myAccount.detectCost(2)
	if myAccount.cost != 2 {
		t.Error("test fee account fail")
	}

	t.Log("Test_FeeAccount pass")
}

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
	prices := []int{1, 3, 2, 1, 8, 4, 9}
	fee := 2
	myProfit := maxProfit(prices, fee)
	targetProfit := 8

	if myProfit != targetProfit {
		t.Error("get max profit fee case fail")
	}

	t.Log("Test_maxProfitFeeCase pass")
}

func Test_maxProfitFeeEmpty(t *testing.T) {
	prices := []int{1}
	fee := 2
	myProfit := maxProfit(prices, fee)
	targetProfit := 0

	if myProfit != targetProfit {
		t.Error("get max profit empty fail")
	}

	t.Log("Test_maxProfitFeeEmpty pass")
}
