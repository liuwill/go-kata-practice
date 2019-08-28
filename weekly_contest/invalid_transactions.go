package weekly_contest

import (
	"math"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	MAX_AMOUNT := 1000
	MAX_MINUTE := 60
	nameDict := map[string]map[int][]string{}

	size := 0
	mark := make([]int, len(transactions))
	for i, v := range transactions {
		meta := strings.Split(v, ",")

		minute, _ := strconv.Atoi(meta[1])
		if num, _ := strconv.Atoi(meta[2]); num > MAX_AMOUNT {
			mark[i] = 1
			size++
		} else if _, ok := nameDict[meta[0]]; ok {
			for curMin, curData := range nameDict[meta[0]] {
				city := curData[1]
				pos, _ := strconv.Atoi(curData[0])
				if city != meta[3] && int(math.Abs(float64(minute-curMin))) <= MAX_MINUTE {
					mark[i] = 1
					size++
				}

				if mark[pos] == 0 {
					mark[pos] = 1
					size++
				}
			}
		}

		if _, ok := nameDict[meta[0]]; !ok {
			nameDict[meta[0]] = map[int][]string{}
		}
		nameDict[meta[0]][minute] = []string{string(i), meta[3]}
	}

	j := 0
	target := make([]string, size)
	for i, v := range mark {
		if v == 1 {
			target[j] = transactions[i]
			j++
		}
	}
	return target
}
