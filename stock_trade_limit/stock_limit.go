package stock_trade_limit

type Trade struct {
	buy    int
	sell   int
	profit int
}

type Account struct {
	cost       int
	profit     int
	hold       bool
	bigTrade   Trade
	smallTrade Trade
}

func (account *Account) buy(price int) {
	account.hold = true
	account.cost = price
}

func (account Account) isHold() bool {
	return account.hold
}

func (account Account) getProfit() int {
	return account.bigTrade.profit + account.smallTrade.profit
}

func (account *Account) sell(price int) {
	account.hold = false
	profit := price - account.cost
	account.addTrade(account.cost, price, profit)

	account.profit = account.profit + profit
	account.cost = 0
}

func (account *Account) addTrade(buyPrice int, sellPrice int, profit int) {
	bigTrade := account.bigTrade
	smallTrade := account.smallTrade

	if profit > bigTrade.profit {
		account.smallTrade = bigTrade
		account.bigTrade = Trade{
			buy:    buyPrice,
			sell:   sellPrice,
			profit: profit,
		}
	} else if profit > smallTrade.profit {
		account.smallTrade = Trade{
			buy:    buyPrice,
			sell:   sellPrice,
			profit: profit,
		}
	}
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	myAccount := Account{
		cost:       0,
		profit:     0,
		hold:       false,
		bigTrade:   Trade{profit: 0},
		smallTrade: Trade{profit: 0},
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
