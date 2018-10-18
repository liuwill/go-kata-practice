package fizzbuzz

import (
	"testing"
)

func Test_FizzBuzz(t *testing.T) {
	expect := 15
	target := fizzBuzz(expect)

	if expect != len(target) {
		t.Error("Run Test_FizzBuzz Fail", target)
	}

	t.Log("Run Test_FizzBuzz Success")
}
