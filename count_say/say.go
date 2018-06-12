package count_say

import "strings"

func say(letter string, count int) string {
	return strings.Repeat(letter, count)
}

func read(input string) string {
	last := ""
	result := ""
	count := 0

	for index, v := range input {
		letter := string(v)
		if index == 0 {
			count = 1
			last = letter
			continue
		}

		if last == letter {
			count++
		} else if last != letter {
			result += say(last, count)
			count = 1
			last = letter
		}
	}
	result += say(last, count)

	return result
}

func countAndSay(n int) string {
	result := "1"
	for i := 2; i <= n; i++ {
		result = read(result)
	}

	return result
}
