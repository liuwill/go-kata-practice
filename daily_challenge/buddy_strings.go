package daily_challenge

import "strings"

/**
 * daily-challenge-859
 * PUZZLE: Buddy Strings
 */
func buddyStringsSwap(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	for i, _ := range A {
		for j := i + 1; j < len(A); j++ {
			newA := strings.Join([]string{A[:i], string(A[j]), A[i+1 : j], string(A[i]), A[j+1:]}, "")
			if newA == B {
				return true
			}
		}
	}
	return false
}

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) == 0 {
		return false
	}

	pos := []int{}
	diff := 0

	hasSame := false
	sameCount := make([]rune, 128)
	for i, v := range A {
		if A[i] != B[i] {
			pos = append(pos, i)
			diff++
		}

		if sameCount[v] > 0 {
			hasSame = true
		}
		sameCount[v]++
	}

	if diff == 0 && hasSame {
		return true
	} else if diff == 2 {
		if A[pos[0]] == B[pos[1]] && B[pos[0]] == A[pos[1]] {
			return true
		}
	}
	return false
}
