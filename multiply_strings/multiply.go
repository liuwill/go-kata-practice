package multiply_strings

import (
	"math"
	"strconv"
	"strings"
)

func byteToInt(input byte) float64 {
	return float64(input - 48)
}

func bitPlus(one byte, two byte) int {
	return int((one - 48) + (two - 48))
}

func shift(num string, bit int) string {
	if bit < 1 {
		return num
	}
	zeros := strings.Repeat("0", bit)
	return strings.Join([]string{num, zeros}, "")
}

func multiply(num1 string, num2 string) string {
	num1List := []byte{}
	num2List := []byte{}

	num1List = append(num1List, num1...)
	num2List = append(num2List, num2...)

	result := 0
	for i := 0; i < len(num1List); i++ {
		pos1 := len(num1List) - 1 - i
		base1 := math.Pow10(i)
		operator1 := byteToInt(num1List[pos1])
		stepResult := 0
		for j := 0; j < len(num2List); j++ {
			pos2 := len(num2List) - 1 - j
			base2 := math.Pow10(j)
			operator2 := byteToInt(num2List[pos2])

			multiResult := int(operator2 * operator1 * base2)
			stepResult = stepResult + multiResult
		}

		result += stepResult * int(base1)
	}

	return strconv.Itoa(result)
}
