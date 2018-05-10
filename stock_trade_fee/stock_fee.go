//best_time_to_buy_and_sell_stock_II
package stock_trade_fee

type Account struct {
	fee    int
	cost   int
	profit int
	hold   bool
}

func (account *Account) detectCost(price int) {
	if account.hold && price < account.cost {
		account.cost = price
	} else if !account.hold {
		account.cost = price
		account.hold = true
	}
}

func (account Account) checkTrade(price int) bool {
	if !account.isHold() {
		return false
	} else if account.cost+account.fee >= price {
		return false
	}

	return true
}

func (account *Account) trade(price int) {
	account.profit = account.profit + price - account.cost - account.fee
	account.hold = false
}

func (account Account) isHold() bool {
	return account.hold
}

func (account Account) getProfit() int {
	return account.profit
}

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	myAccount := Account{
		fee:    fee,
		cost:   0,
		profit: 0,
		hold:   false,
	}

	for _, current := range prices {
		if !myAccount.checkTrade(current) {
			myAccount.detectCost(current)
		} else {
			myAccount.trade(current)
		}
	}

	return myAccount.getProfit()
}
