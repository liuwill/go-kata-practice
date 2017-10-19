package fizzbuzz

import (
  "testing"
  "strconv"
)

func Test_TranslateFirst(t *testing.T) {
  if translateFirst(1) != "1" {
    t.Error("translateFirst fail")
  }

  if translateFirst(3) != "Fizz" {
    t.Error("translateFirst fail")
  }

  if translateFirst(5) != "Buzz" {
    t.Error("translateFirst fail")
  }
  t.Log("I'm pass")
}

func Test_Run(t *testing.T) {
  expectMap := make(map[string]string)
  expectMap["1"] = "1"
  expectMap["2"] = "2"
  expectMap["3"] = "Fizz"
  expectMap["5"] = "Buzz"
  expectMap["6"] = "Fizz"
  expectMap["10"] = "Buzz"
  expectMap["15"] = "FizzBuzz"
  expectMap["30"] = "FizzBuzz"

  for key, val := range expectMap {
    num, _ := strconv.Atoi(key)
    if translate(num) != val {
      t.Error("translate fail")
    }
  }

    t.Log("Translate Succeed")
}
