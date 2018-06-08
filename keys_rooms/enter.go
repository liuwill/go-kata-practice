package keys_rooms

type Room struct {
	Visited bool
	HasKey  bool
}

func canVisitAllRooms(rooms [][]int) bool {
	roomMap := make([]Room, len(rooms))
	keys := []int{}
	visited := 1

	initkeys := rooms[0]
	keys = append(keys, initkeys...)
	roomMap[0] = Room{
		Visited: true,
		HasKey:  true,
	}

	for len(keys) > 0 && visited < len(rooms) {
		index := keys[0]
		keys = keys[1:]

		current := roomMap[index]
		roomMap[index].HasKey = true
		if current.Visited {
			continue
		}

		roomMap[index].Visited = true
		currentKeys := rooms[index]
		for _, key := range currentKeys {
			keyRoom := roomMap[key]
			if keyRoom.Visited || keyRoom.HasKey {
				continue
			}
			roomMap[key].HasKey = true
			keys = append(keys, key)
		}
	}

	return visited == len(rooms)
}
