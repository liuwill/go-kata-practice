package dungeon_game

type Action struct {
	X       int
	Y       int
	Point   int
	Energy  int
	Consume int
}

/*
func (action *Action) isBetterAction(x int, y int, marks [][]int) bool {
	if marks[x][y] == 0 || marks[x][y] > action.Energy {
		return true
	}
	return false
}
*/

func (action *Action) countNextAction(x int, y int, dungeon [][]int, marks [][]int) Action {
	currentState := dungeon[x][y]
	nextEnergy := action.Energy - currentState
	marks[x][y] = nextEnergy
	currentConsume := action.Consume
	if currentConsume < nextEnergy {
		currentConsume = nextEnergy
	}

	return Action{
		X:       x,
		Y:       y,
		Point:   dungeon[x][y],
		Consume: currentConsume,
		Energy:  nextEnergy,
	}
}

func (action *Action) nextActions(width int, height int, dungeon [][]int, marks [][]int) []Action {
	target := []Action{}
	if action.X+1 < width { // && action.isBetterAction(action.X+1, action.Y, marks)
		xAction := action.countNextAction(action.X+1, action.Y, dungeon, marks)
		target = append(target, xAction)
	}
	if action.Y+1 < height { // && action.isBetterAction(action.X, action.Y+1, marks)
		yAction := action.countNextAction(action.X, action.Y+1, dungeon, marks)
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

func (action *Action) isBest(best int) bool {
	if action.Consume > 0 && best == 0 {
		return true
	}

	return action.Consume < best
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
	initAction := Action{
		X:       0,
		Y:       0,
		Point:   dungeon[0][0],
		Consume: 0,
		Energy:  initEnergy,
	}

	if initEnergy > 0 {
		initAction.Consume = initEnergy
	}
	actions := []Action{initAction}

	best := 0
	for len(actions) > 0 {
		currentAction := actions[0]
		if currentAction.isEnd(width, height) && currentAction.isBest(best) {
			// fmt.Printf("%v ==== \n", currentAction)
			best = currentAction.Consume
		}

		nextActions := currentAction.nextActions(width, height, dungeon, marks)
		// fmt.Printf("%v %v\n", currentAction, (nextActions))
		actions = append(actions[1:], nextActions...)
	}

	return best + 1
}
