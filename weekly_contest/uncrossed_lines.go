package weekly_contest

func maxUncrossedLines(A []int, B []int) int {
	aLen := len(A)
	bLen := len(B)
	dict := make([][]int, aLen)
	for i := 0; i < len(dict); i++ {
		dict[i] = make([]int, bLen)
	}

	for i := 0; i < aLen; i++ {
		if A[i] == B[0] {
			for k := i; k < aLen; k++ {
				dict[k][0] = 1
			}
			break
		}
	}

	for j := 0; j < bLen; j++ {
		if A[0] == B[j] {
			for k := j; k < bLen; k++ {
				dict[0][k] = 1
			}
			break
		}
	}

	for i := 1; i < aLen; i++ {
		for j := 1; j < bLen; j++ {
			if A[i] == B[j] {
				dict[i][j] = dict[i-1][j-1] + 1
			} else {
				dict[i][j] = dict[i-1][j]
				if dict[i][j-1] > dict[i-1][j] {
					dict[i][j] = dict[i][j-1]
				}
			}
		}
	}

	return dict[aLen-1][bLen-1]
}
