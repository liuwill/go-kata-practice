package daily_challenge

import "strings"

/**
 * daily-challenge-459
 * PUZZLE: Repeated Substring Pattern
 */
func repeatedSubstringPattern(s string) bool {
	ll := len(s)
	for m := ll / 2; m > 0; m-- {
		if ll%m != 0 {
			continue
		}
		pass := true

		sub := s[:m]
		for i := 0; i < ll; i += m {
			cur := s[i : m+i]
			if sub != cur {
				pass = false
				break
			}
		}

		if pass {
			return true
		}
	}
	return false
}

func repeatedSubstringPatternFast(s string) bool {
	ll := len(s)
	if ll == 0 {
		return false
	}

	chk := (s + s)[1 : ll*2-1]
	return strings.Contains(chk, s)
}
