package weekly_contest_93

type MarkItem struct {
	Number   int
	Position int
}

func sortMark(list []int) []MarkItem {
	markList := make([]MarkItem, len(list))
	for i := 0; i < len(list); i++ {
		markList[i] = MarkItem{
			Position: i,
			Number:   list[i],
		}
	}

	for i := 0; i < len(markList); i++ {
		for j := i + 1; j < len(markList); j++ {
			if markList[j].Number < markList[i].Number {
				temp := markList[i]
				markList[i] = markList[j]
				markList[j] = temp
			}
		}
	}

	return markList
}

func advantageCount(A []int, B []int) []int {
	sortedB := sortMark(B)
	sortedA := sortMark(A)
	target := make([]int, len(A))
	markedA := make([]bool, len(A))

	for _, value := range sortedB {
		winner := false
		for i := 0; i < len(sortedA); i++ {
			pos := sortedA[i].Position
			if sortedA[i].Number > value.Number && !markedA[pos] {
				markedA[pos] = true
				target[value.Position] = A[pos]
				winner = true
				break
			}
		}

		if !winner {
			for i := 0; i < len(sortedA); i++ {
				pos := sortedA[i].Position
				if !markedA[pos] {
					markedA[pos] = true
					target[value.Position] = A[pos]
					winner = true
					break
				}
			}
		}

	}

	return target
}
