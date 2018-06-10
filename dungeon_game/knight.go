package dungeon_game

type Action struct {
	X       int
	Y       int
	Point   int
	Energy  int
	Consume int
}

func (action *Action) countNextAction(x int, y int, dungeon [][]int) Action {
	currentState := dungeon[x][y]
	nextEnergy := action.Energy - currentState

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

func (action *Action) nextActions(width int, height int, dungeon [][]int) []Action {
	target := []Action{}
	if action.X+1 < width { // && action.isBetterAction(action.X+1, action.Y, marks)
		xAction := action.countNextAction(action.X+1, action.Y, dungeon)
		target = append(target, xAction)
	}
	if action.Y+1 < height { // && action.isBetterAction(action.X, action.Y+1, marks)
		yAction := action.countNextAction(action.X, action.Y+1, dungeon)
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
	if action.Consume > 0 && best <= 0 {
		return true
	}

	return action.Consume < best
}

func calculateKnightMinimumHP(dungeon [][]int) int {
	width := len(dungeon)
	height := 0
	if len(dungeon) > 0 {
		height = len(dungeon[0])
	}

	initEnergy := 1 - dungeon[0][0]
	initAction := Action{
		X:       0,
		Y:       0,
		Point:   dungeon[0][0],
		Consume: 1,
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
			best = currentAction.Consume
		}

		actions = actions[1:]
		nextActions := currentAction.nextActions(width, height, dungeon)
		// fmt.Printf("%v %v\n", currentAction, (nextActions))
		for _, next := range nextActions {
			if next.Consume < best || best <= 0 {
				actions = append(actions, next)
			}
		}
	}

	return best
}
