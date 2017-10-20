package fizzbuzz_plus

import (
  "strconv"
  "strings"
)

func translate(num int) string {
  if num % 5 == 0 || strings.Contains(strconv.Itoa(num), "5") {
    return "Buzz"
  }
  if num % 3 == 0 || strings.Contains(strconv.Itoa(num), "3") {
    return "Fizz"
  }

  return strconv.Itoa(num)
}
