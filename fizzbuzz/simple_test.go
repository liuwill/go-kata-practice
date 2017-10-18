package fizzbuzz

import (
  "testing"
)

func Test_Run(t *testing.T) {
  if translate(1) != "1" {
    t.Log(translate(1))
    t.Error("tranlate fail")
  }

  if translate(3) != "Fizz" {
    t.Error("translate fail")
  }

  if translate(5) != "Buzz" {
    t.Error("translate fail")
  }
  t.Log("I'm pass")
}
