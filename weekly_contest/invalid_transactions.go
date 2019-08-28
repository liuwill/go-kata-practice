package weekly_contest

import (
	"math"
	"strconv"
	"strings"
)

const (
	MAX_AMOUNT = 1000
	MAX_MINUTE = 60
)

type Meta struct {
	Pos  int
	City string
}

/**
 * daily-challenge-1169
 * PUZZLE: Invalid Transactions
 */
func invalidTransactions(transactions []string) []string {
	nameDict := map[string]map[int]Meta{}

	mark := make([]int, len(transactions))
	for i, v := range transactions {
		meta := strings.Split(v, ",")

		minute, _ := strconv.Atoi(meta[1])
		if num, _ := strconv.Atoi(meta[2]); num > MAX_AMOUNT {
			mark[i] = 1
		}

		if _, ok := nameDict[meta[0]]; ok {
			for curMin, curMeta := range nameDict[meta[0]] {
				if curMeta.City != meta[3] && int(math.Abs(float64(minute-curMin))) <= MAX_MINUTE {
					mark[i] = 1
					mark[curMeta.Pos] = 1
				}
			}
		}

		if _, ok := nameDict[meta[0]]; !ok {
			nameDict[meta[0]] = map[int]Meta{}
		}
		nameDict[meta[0]][minute] = Meta{
			City: meta[3],
			Pos:  i,
		}
	}

	target := []string{}
	for i, v := range mark {
		if v == 0 {
			continue
		}

		target = append(target, transactions[i])
	}
	return target
}
