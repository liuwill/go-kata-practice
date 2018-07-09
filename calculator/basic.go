package calculator

import (
	"strconv"
)

func isNumber(letter rune) bool {
	return letter >= 48 && letter <= 57
}

func isHighLevelOperator(letter string) bool {
	return letter == "*" || letter == "/"
}

func isLowLevelOperator(letter string) bool {
	return letter == "+" || letter == "-"
}

func operate(operator string, left string, right string) string {
	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)

	result := 0
	if operator == "+" {
		result = leftNum + rightNum
	} else if operator == "-" {
		result = leftNum - rightNum
	} else if operator == "*" {
		result = leftNum * rightNum
	} else if operator == "/" {
		result = leftNum / rightNum
	}
	return strconv.Itoa(result)
}

func calculate(s string) int {
	current := ""
	operator := ""
	numbers := []string{}
	operators := []string{}
	for _, v := range s {
		letter := string(v)
		if letter == " " {
			continue
		}

		if isNumber(v) {
			current += letter
		} else if len(operator) > 0 {
			numbers[len(numbers)-1] = operate(operator, numbers[len(numbers)-1], current)
			operator = ""
			current = ""
		} else {
			numbers = append(numbers, current)
			current = ""
		}

		if isHighLevelOperator(letter) {
			operator = letter
		} else if isLowLevelOperator(letter) {
			operators = append(operators, letter)
		}
	}
	if len(current) > 0 {
		if len(operator) > 0 {
			numbers[len(numbers)-1] = operate(operator, numbers[len(numbers)-1], current)
			operator = ""
			current = ""
		} else {
			numbers = append(numbers, current)
			current = ""
		}
	}

	result := numbers[0]
	for i := 0; i < len(operators); i++ {
		operator = operators[i]

		result = operate(operator, result, numbers[i+1])
	}
	target, _ := strconv.Atoi(result)
	return target
}
