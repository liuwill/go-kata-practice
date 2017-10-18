package fizzbuzz

import (
  "strconv"
)

func translate(num int) string {
  if num % 3 == 0 {
    return "Fizz"
  } else if num % 5 == 0 {
    return "Buzz"
  }

  return strconv.Itoa(num)
}
