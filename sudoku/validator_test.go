package sudoku

import (
	"testing"
)

func Test_CheckLine_PASS(t *testing.T) {
	source := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isLineValidate := checkLine(source)
	if !isLineValidate {
		t.Error("line check fail")
	}

	t.Log("Test_CheckLine pass")
}

func Test_CheckLine_FAIL(t *testing.T) {
	source := [][]byte{
		{'8', '3', '.', '.', '7', '.', '3', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	}
	isLineValidate := checkLine(source)
	if isLineValidate {
		t.Error("line check fail")
	}

	t.Log("Test_CheckLine pass")
}
