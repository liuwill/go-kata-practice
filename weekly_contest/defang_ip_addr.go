package weekly_contest

/**
 * daily-challenge-1108
 * PUZZLE: Defanging an IP Address
 */
func defangIPaddr(address string) string {
	result := ""

	for _, letter := range address {
		if letter == '.' {
			result += "[.]"
		} else {
			result += string(letter)
		}
	}
	return result
}
