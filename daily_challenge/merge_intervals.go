package daily_challenge

import "sort"

/**
 * daily-challenge-56
 * PUZZLE: Merge Intervals
 */

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
	output := []Interval{}

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	var current Interval = intervals[0]
	for i := 1; i < len(intervals); i++ {
		if current.End < intervals[i].Start {
			output = append(output, current)
			current = intervals[i]
		} else {
			start := current.Start
			end := intervals[i].End
			// if current.Start > intervals[i].Start {
			// 	start = intervals[i].Start
			// }
			if intervals[i].End < current.End {
				end = current.End
			}

			current = Interval{
				Start: start,
				End:   end,
			}
		}
	}
	output = append(output, current)
	return output
}
