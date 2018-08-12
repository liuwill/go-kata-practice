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
	return target
}
