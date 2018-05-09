package stock_trade_limit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := 0
	buyPrice := 0
	for i := len(prices) - 1; i >= 0; i-- {
		current := prices[i]

		estimate := buyPrice - current
		if buyPrice < current && i > 0 {
			buyPrice = current
		} else if estimate > profit {
			profit = estimate
		}
	}

	return profit
}
