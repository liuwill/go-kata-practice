package weekly_contest

const MAX_VOWEL_COUNT = 1000000007

func initVowelCounter() map[rune]int {
	return map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}
}

func countVowelPermutation(n int) int {
	current := map[rune]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}
	addFuncs := map[rune]func(map[rune]int, int) int{
		'a': func(dict map[rune]int, count int) int {
			dict['e'] += count

			return count
		},
		'e': func(dict map[rune]int, count int) int {
			dict['a'] += count
			dict['i'] += count

			return count * 2
		},
		'i': func(dict map[rune]int, count int) int {
			dict['a'] += count
			dict['e'] += count
			dict['o'] += count
			dict['u'] += count

			return count * 4
		},
		'o': func(dict map[rune]int, count int) int {
			dict['i'] += count
			dict['u'] += count

			return count * 2
		},
		'u': func(dict map[rune]int, count int) int {
			dict['a'] += count

			return count
		},
	}

	for j := 2; j <= n; j++ {
		round := 0
		newCounter := initVowelCounter()
		for k, v := range current {
			round += addFuncs[k](newCounter, v)
		}

		if round >= MAX_VOWEL_COUNT {
			return MAX_VOWEL_COUNT
		}
		current = newCounter
	}

	result := 0
	for _, v := range current {
		result += v
	}
	return result
}
