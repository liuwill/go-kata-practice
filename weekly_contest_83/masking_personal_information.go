package weekly_contest_83

import "strings"

func isEmail(input string) bool {
	return strings.Contains(input, "@")
}

func maskEmail(email string) string {
	meta := strings.Split(email, "@")
	name := string(meta[0][0])
	name += "*****"
	name += string(meta[0][len(meta[0])-1])
	return name + "@" + meta[1]
}

func maskPhone(phone string) string {
	numbers := ""
	for _, v := range phone {
		if v >= 48 && v <= 57 {
			numbers += string(v)
		}
	}

	index := len(numbers) - 4
	result := numbers[index:]
	length := index
	for length > 0 {
		maskLen := length
		if length > 3 {
			maskLen = 3
			length -= 3
		} else {
			length = 0
		}

		if maskLen > 0 {
			maskRepeat := strings.Repeat("*", maskLen)
			result = maskRepeat + "-" + result
		}
	}

	if len(numbers) > 10 {
		result = "+" + result
	}
	return result
}

func maskPII(S string) string {
	if isEmail(S) {
		return maskEmail(strings.ToLower(S))
	}

	return maskPhone(S)
}
