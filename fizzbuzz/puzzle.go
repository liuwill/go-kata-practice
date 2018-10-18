package fizzbuzz

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		current := ""

		if i%3 == 0 {
			current += "Fizz"
		}

		if i%5 == 0 {
			current += "Buzz"
		}

		if len(current) < 1 {
			current += strconv.Itoa(i)
		}

		result[i-1] = current
	}
	return result
}

func fizzBuzzForSpeed(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		current := ""

		isFizz := i%3 == 0
		isBuzz := i%5 == 0
		if isFizz && isBuzz {
			current = "FizzBuzz"
		} else if isFizz {
			current = "Fizz"
		} else if isBuzz {
			current = "Buzz"
		} else {
			current = strconv.Itoa(i)
		}

		// println(current, i)
		result[i-1] = current
	}
	return result
}
