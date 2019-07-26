package weekly_contest

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
