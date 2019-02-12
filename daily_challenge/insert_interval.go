package daily_challenge

/**
 * daily-challenge-57
 * PUZZLE: Insert Interval
 */
type Interval struct {
	Start int
	End   int
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func insert(intervals []Interval, newInterval Interval) []Interval {
	output := []Interval{}
	scopeStart := -1
	scopeEnd := -1
	current := 0
	for i, item := range intervals {
		if item.Start > newInterval.End {
			break
		}

		current = i
		if item.End < newInterval.Start {
			output = append(output, item)
			continue
		}

		if scopeStart == -1 {
			scopeStart = i
		}
		scopeEnd = i
	}

	for i := scopeStart; i < scopeEnd; i++ {

	}
	return intervals
}
