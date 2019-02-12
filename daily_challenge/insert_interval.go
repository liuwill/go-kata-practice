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

	start := newInterval.Start
	end := newInterval.End
	if scopeStart > -1 && intervals[scopeStart].Start < start {
		start = intervals[scopeStart].Start
	}

	if scopeEnd > -1 && intervals[scopeEnd].End > end {
		end = intervals[scopeEnd].End
	}
	output = append(output, Interval{
		Start: start,
		End:   end,
	})

	for i := current + 1; i < len(intervals); i++ {
		output = append(output, intervals[i])
	}
	// println(scopeStart, scopeEnd)
	// fmt.Printf("Intervals: %v - %v", output, newInterval)
	return output
}
