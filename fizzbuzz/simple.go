package fizzbuzz

import (
  "strconv"
)

func translateFirst(num int) string {
  if num % 3 == 0 {
    return "Fizz"
  } else if num % 5 == 0 {
    return "Buzz"
  }

  return strconv.Itoa(num)
}

func translate(num int) string {
  result := ""
  if isFizz(num) {
    result = "Fizz"
  }
  if isBuzz(num) {
    result += "Buzz"
  }
  if len(result) == 0 {
    return strconv.Itoa(num)
  }

  return result
}

func isFizz(num int) bool {
  return num % 3 == 0
}

func isBuzz(num int) bool {
  return num % 5 == 0
}
