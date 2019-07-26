package weekly_contest

/**
 * daily-challenge-1109
 * PUZZLE: Corporate Flight Bookings
 */
func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)

	for i := 0; i < len(bookings); i++ {
		booking := bookings[i]
		for j := booking[0] - 1; j < booking[1]; j++ {
			result[j] += booking[2]
		}
	}
	return result
}
