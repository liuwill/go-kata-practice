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

func (account *Account) sell(price int) {
	account.hold = false
	profit := price - account.cost
	account.profit = account.profit + profit
	account.cost = 0
}

func maxProfit(prices []int) int {
	var result int
	myAccount := Account{
		cost:   0,
		profit: 0,
		hold:   false,
	}

	current := prices[0]
	for i := 1; i < len(prices); i++ {
		if !myAccount.isHold() && current < prices[i] {
			myAccount.buy(prices[i])
		} else if myAccount.isHold() {
			if current > prices[i] || (i == len(prices)-1 && prices[i] > myAccount.cost) {
				myAccount.sell(prices[i])
			}
		}
		current = prices[i]
	}

	return result
}
