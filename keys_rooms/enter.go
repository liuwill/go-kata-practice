package keys_rooms

type Room struct {
	Visited bool
	HasKey  bool
}

func canVisitAllRooms(rooms [][]int) bool {
	// roomMap := make([]Room, len(rooms))
	keys := []int{}
	visited := 0

	initkeys := rooms[0]
	keys = append(keys, initkeys...)
	for len(keys) > 0 && visited < len(rooms) {

	}

	return visited == len(rooms)
}
