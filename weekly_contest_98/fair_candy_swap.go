package weekly_contest_98

func sumCandy(list []int) int {
	sum := 0
	for _, value := range list {
		sum += value
	}

	return sum
}

func isSwapPass(currentA int, currentB int, countA int, countB int) bool {
	if countA > countB && currentA > currentB {
		return false
	}

	if countA < countB && currentA < currentB {
		return false
	}

	if countA-currentA+currentB == countB-currentB+currentA {
		return true
	}

	return false
}

func fairCandySwap(A []int, B []int) []int {
	countA := sumCandy(A)
	countB := sumCandy(B)

	result := []int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if isSwapPass(A[i], B[j], countA, countB) {
				result[0] = A[i]
				result[1] = B[j]
				break
			}
		}
	}
	return result
}
