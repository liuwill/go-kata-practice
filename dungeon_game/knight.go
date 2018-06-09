package dungeon_game

type Action struct {
	X       int
	Y       int
	Consume int
	Energy  int
}

/*
func (action *Action) isBetterAction(x int, y int, marks [][]int) bool {
	if marks[x][y] == 0 || marks[x][y] > action.Energy {
		return true
	}
	return false
}
*/

func (action *Action) countNextAction(x int, y int, dungeon [][]int) Action {
	currentState := dungeon[x][y]
	nextEnergy := action.Energy - currentState
	return Action{
		X:      x,
		Y:      y,
		Energy: nextEnergy,
	}
}

func (action *Action) nextActions(width int, height int, dungeon [][]int, marks [][]int) []Action {
	target := []Action{}
	if action.X+1 < width { // && action.isBetterAction(action.X+1, action.Y, marks)
		xAction := action.countNextAction(action.X+1, action.Y, dungeon)
		marks[action.X+1][action.Y] = xAction.Energy
		target = append(target, xAction)
	}
	if action.Y+1 < height { // && action.isBetterAction(action.X, action.Y+1, marks)
		yAction := action.countNextAction(action.X, action.Y+1, dungeon)
		marks[action.X][action.Y] = yAction.Energy
		target = append(target, yAction)
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
	height := len(dungeon)
	width := 0
	if len(dungeon) > 0 {
		width = len(dungeon[0])
	}

	initEnergy := -dungeon[0][0]
	actions := []Action{Action{
		X:      0,
		Y:      0,
		Energy: initEnergy,
	}}

	best := 0
	for len(actions) > 0 {
		currentAction := actions[0]
		if currentAction.isEnd(width, height) && currentAction.Energy <= best {
			best = currentAction.Energy
		}

		// fmt.Printf("%v ==== \n", currentAction)
		nextActions := currentAction.nextActions(width, height, dungeon, marks)
		// fmt.Printf("%v %v\n", currentAction, (nextActions))
		actions = append(actions[1:], nextActions...)
	}

	return best
}
