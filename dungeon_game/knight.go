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

func (action *Action) countEnergy(x int, y int, dungeon [][]int) int {
	currentState := dungeon[x][y]
	return action.Energy - currentState
}

func (action *Action) nextActions(width int, height int, dungeon [][]int, marks [][]int) []Action {
	target := []Action{}
	if action.X+1 < width && action.isBetterAction(action.X+1, action.Y, marks) {
		xEnergy := action.countEnergy(action.X+1, action.Y, dungeon)
		marks[action.X+1][action.Y] = xEnergy
		target = append(target, Action{
			X:      action.X + 1,
			Y:      action.Y,
			Energy: xEnergy,
		})
	}
	if action.Y+1 < height && action.isBetterAction(action.X, action.Y+1, marks) {
		yEnergy := action.countEnergy(action.X, action.Y+1, dungeon)
		marks[action.X+1][action.Y] = yEnergy
		target = append(target, Action{
			X:      action.X,
			Y:      action.Y + 1,
			Energy: yEnergy,
		})
	}
	return target
}

func (action *Action) isEnd(width int, height int) bool {
	if action.X == width-1 && action.Y == height-1 {
		return true
	}
	return false
}

func calculateMinimumHP(dungeon [][]int) int {
	marks := make([][]int, len(dungeon))
	for index, line := range dungeon {
		marks[index] = make([]int, len(line))
	}
	width := len(dungeon)
	height := 0
	if len(dungeon) > 0 {
		height = len(dungeon[0])
	}

	actions := []Action{Action{
		X:      0,
		Y:      0,
		Energy: -dungeon[0][0],
	}}
	best := 0
	for len(actions) > 0 {
		currentAction := actions[0]
		if currentAction.isEnd(width, height) && currentAction.Energy <= best {
			best = currentAction.Energy
			continue
		}

		currentAction.nextActions(width, height, dungeon, marks)
	}

	return best
}
