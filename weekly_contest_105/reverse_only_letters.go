package weekly_contest_105

func isLetter(letter byte) bool {
	if letter >= 97 && letter <= 122 {
		return true
	} else if letter >= 65 && letter <= 90 {
		return true
	}

	return false
}

func reverseOnlyLetters(S string) string {
	result := ""
	j := len(S) - 1
	for i := 0; i < len(S); {
		if !isLetter(S[i]) {
			result += string(S[i])
			i++
		} else if j >= 0 {
			if !isLetter(S[j]) {
				j--
			} else {
				result += string(S[j])
				i++
				j--
			}
		}
	}

	return result
}
