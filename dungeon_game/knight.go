package dungeon_game

type Action struct {
	X      int
	Y      int
	Energy int
}

func (action *Action) isBetterAction(x int, y int, marks [][]int) bool {
	if marks[x][y] == 0 || marks[x][y] > action.Energy {
		return true
	}
	return false
}

func countEnergy(point Action, x int, y int, dungeon [][]int, marks [][]int) {

}

func (action *Action) nextActions(width int, height int, dungeon [][]int, marks [][]int) []Action {
	target := []Action{}
	if action.X+1 < width && action.isBetterAction(action.X+1, action.Y, marks) {
		target = append(target, Action{
			X: action.X + 1,
			Y: action.Y,
		})
	}
	if action.Y+1 < height && action.isBetterAction(action.X, action.Y+1, marks) {
		target = append(target, Action{
			X: action.X,
			Y: action.Y + 1,
		})
	}
	return target
}

func calculateMinimumHP(dungeon [][]int) int {
	pointMap := make([][]int, len(dungeon))
	for index, line := range dungeon {
		pointMap[index] = make([]int, len(line))
	}

	return 0
}
