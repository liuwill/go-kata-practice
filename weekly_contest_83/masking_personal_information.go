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
	return phone
}

func maskPII(S string) string {
	if isEmail(S) {
		return maskEmail(strings.ToLower(S))
	}

	return maskPhone(S)
}
