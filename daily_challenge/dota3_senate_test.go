package daily_challenge

import "testing"

func Test_PredictPartyVictory(t *testing.T) {
	source := "RD"
	target := predictPartyVictory(source)
	expect := Radiant

	if expect != target {
		t.Error("Translate Test_PredictPartyVictory Fail", target)
	}
	t.Log("Translate Test_PredictPartyVictory Success")
}

func Test_PredictPartyVictoryDire(t *testing.T) {
	source := "RDD"
	target := predictPartyVictory(source)
	expect := Dire

	if expect != target {
		t.Error("Translate Test_PredictPartyVictoryDire Fail", target)
	}
	t.Log("Translate Test_PredictPartyVictoryDire Success")
}

func Test_PredictPartyVictoryFail(t *testing.T) {
	source := "DDRRR"
	target := predictPartyVictory(source)
	expect := Dire

	if expect != target {
		t.Error("Translate Test_PredictPartyVictoryFail Fail", target)
	}
	t.Log("Translate Test_PredictPartyVictoryFail Success")
}

func Test_PredictPartyVictoryLong(t *testing.T) {
	source := "DDDDRRDDDRDRDRRDDRDDDRDRRRRDRRRRRDRDDRDDRRDDRRRDDRRRDDDDRRRRRRRDDRRRDDRDDDRRRDRDDRDDDRRDRRDRRRDRDRDR"
	target := predictPartyVictory(source)
	expect := Dire

	if expect != target {
		t.Error("Translate Test_PredictPartyVictoryLong Fail", target)
	}
	t.Log("Translate Test_PredictPartyVictoryLong Success")
}
