package daily_challenge

import "strings"

/**
 * daily-challenge-859
 * PUZZLE: Buddy Strings
 */
func buddyStrings(A string, B string) bool {
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
