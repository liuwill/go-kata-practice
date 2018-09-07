package daily_challenge

import "fmt"

func longestValidParentheses(s string) int {
	println(s, "===========")
	max := 0
	left := '('
	right := ')'

	match := 0
	// listPos := 0
	matchList := make([]int, len(s))
	last := []int{}
	for i, v := range s {
		fmt.Printf("matchList=%v last=%v", matchList, last)
		println(" match=", match, " i=", i, " v=", string(v))
		if v == left {
			matchList[match] = i
			match++
		} else if v == right {
			if match <= 0 {
				match = 0
				continue
			}

			match--
			current := i - matchList[match] + 1
			if len(last) > 0 && matchList[match]-1 == last[1] {
				last = []int{last[0], i}
				current = i - last[0] + 1
			} else {
				last = []int{matchList[match], i}
			}

			if current > max {
				max = current
			}
		}
	}

	return max
}
