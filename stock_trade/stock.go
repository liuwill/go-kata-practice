package stock_trade

type Account struct {
	cost   int
	profit int
	hold   bool
}

func (account *Account) buy(price int) {
	account.hold = true
	account.cost = price
}

func (account Account) isHold() bool {
	return account.hold
}

func (account Account) getProfit() int {
	return account.profit
}

func (account *Account) sell(price int) {
	account.hold = false
	profit := price - account.cost
	account.profit = account.profit + profit
	account.cost = 0
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	myAccount := Account{
		cost:   0,
		profit: 0,
		hold:   false,
	}

	for i, current := range prices {
		nextPos := i + 1
		if !myAccount.isHold() {
			if nextPos < len(prices) && current < prices[nextPos] {
				myAccount.buy(current)
			}
		} else if myAccount.isHold() {
			if i == len(prices)-1 && current > myAccount.cost {
				myAccount.sell(current)
			} else if current > prices[nextPos] {
				myAccount.sell(current)
			}
		}
	}

	return myAccount.getProfit()
}
