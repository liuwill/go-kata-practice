package weekly_contest

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
	if B > A {
		current = 'b'
	}
	size := 0

	max := A + B

	for len(target) < max {
		reverse := reverseDict[current]

		limitSize := MAX_DUPLICATE_SIZE
		if letterDict[reverse] > letterDict[current] {
			limitSize = 1
		}

		if letterDict[current] <= 0 || size >= limitSize {
			current = reverse
			size = 0
		} else {
			size++
			target += string(current)
			letterDict[current]--
		}
	}

	return target
}
