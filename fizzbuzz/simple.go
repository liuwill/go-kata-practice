package fizzbuzz

import (
  "strconv"
)

/**
 * translate First
 */
func translateFirst(num int) string {
  if num % 3 == 0 {
    return "Fizz"
  } else if num % 5 == 0 {
    return "Buzz"
  }

  return strconv.Itoa(num)
}

/**
 * translate Second
 */
func translateFull(num int) string {
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

/**
 * traslate Third
 */
func translateBetter(num int) string {
  result := ""
  handleFizz(num, &result)
  handleBuzz(num, &result)
  handleOther(num, &result)

  return result
}

func handleOther(num int, resultTarget *string) string {
  if len(*resultTarget) == 0 {
    *resultTarget = strconv.Itoa(num)
  }
  return *resultTarget
}

func handleFizz(num int, resultTarget *string) string {
  if isFizz(num) {
    *resultTarget += "Fizz"
  }
  return *resultTarget
}

func handleBuzz(num int, resultTarget *string) string {
  if isBuzz(num) {
    *resultTarget += "Buzz"
  }
  return *resultTarget
}

func isFizz(num int) bool {
  return num % 3 == 0
}

func isBuzz(num int) bool {
  return num % 5 == 0
}

/**
 * Translate Last
 */
func isMonkey(num int) bool {
  return num % 7 == 0
}

func handleMonkey(num int, resultTarget *string) string {
  if isMonkey(num) {
    *resultTarget += "Monkey"
  }
  return *resultTarget
}

func buildTranslate(funcList []func(num int, resultTarget *string) string) func(num int) string {
  return func (num int) string {
    result := ""
    for _, handler := range funcList {
      handler(num, &result)
    }
    return result
  }
}
