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
		println(current, i)
	}
	return result
}
