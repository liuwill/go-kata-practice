package weekly_content

/**
 * weekly-contest-121
 * PUZZLE: String Without AAA or BBB
 */

const MAX_DUPLICATE_SIZE = 2

func strWithout3a3b(A int, B int) string {
	target := ""

	letterDict := map[rune]int{
		'a': A,
		'b': B,
	}

	reverseDict := map[rune]rune{
		'a': 'b',
		'b': 'a',
	}

	current := 'a'
	size := 0

	max := A + B

	for len(target) < max {
		if letterDict[current] < 0 || size >= MAX_DUPLICATE_SIZE {
			current = reverseDict[current]
			size = 0
		} else {
			size++
			target += string(current)
		}
	}

	return target
}
