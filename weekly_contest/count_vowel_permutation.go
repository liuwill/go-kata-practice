package weekly_contest

const MAX_VOWEL_COUNT = 1e9 + 7

func initVowelCounter(initVal int) map[rune]int {
	return map[rune]int{
		'a': initVal,
		'e': initVal,
		'i': initVal,
		'o': initVal,
		'u': initVal,
	}
}

/**
 * weekly-contest-1220
 * PUZZLE: Count Vowels Permutation
 */
func countVowelPermutation(n int) int {
	current := initVowelCounter(1)
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
		newCounter := initVowelCounter(0)
		for k, v := range current {
			// println(k, v, j, n)
			round += addFuncs[k](newCounter, v%MAX_VOWEL_COUNT)
		}

		// if round >= MAX_VOWEL_COUNT {
		// 	return MAX_VOWEL_COUNT
		// }
		current = newCounter
	}

	result := 0
	for _, v := range current {
		result += v
	}
	return result % MAX_VOWEL_COUNT
}
