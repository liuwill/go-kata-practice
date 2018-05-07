package fizzbuzz_plus

import (
  "testing"
  "strconv"
)

func Test_translate(t *testing.T) {
  numberMap := make(map[string]string)
  numberMap["1"] = "1"
  numberMap["3"] = "Fizz"
  numberMap["5"] = "Buzz"
  numberMap["15"] = "Buzz"
  numberMap["31"] = "Fizz"
  numberMap["51"] = "Buzz"

  for num, result := range numberMap {
    theNum, _ := strconv.Atoi(num)
    target := translate(theNum)

    if target != result {
      t.Error("Translate Fail For " + num)
    }
  }
  t.Log("Translate Success")
}
