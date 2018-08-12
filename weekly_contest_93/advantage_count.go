package weekly_contest_93

type MarkItem struct {
	Number int
	Mark   bool
}

func sortMark(list []int) []MarkItem {
	markList := []MarkItem{}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[i] {
				temp := list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}

		markList = append(markList[:], MarkItem{
			Mark:   false,
			Number: list[i],
		})
	}

	return markList
}

func advantageCount(A []int, B []int) []int {
	target := []int{}
	markedA := sortMark(A)

	for _, value := range B {
		winner := false
		for _, markedItem := range markedA {
			if markedItem.Number > value && !markedItem.Mark {
				markedItem.Mark = true
				target = append(target[:], markedItem.Number)
				winner = true
				break
			}
		}

		if !winner {
			for _, markedItem := range markedA {
				if !markedItem.Mark {
					markedItem.Mark = true
					target = append(target[:], markedItem.Number)
					break
				}
			}
		}
	}

	return target
}
