package weekly_contest

import (
	"testing"
)

func Test_CamelMatch(t *testing.T) {
	sourceCase := [][]string{
		[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
	}
	patternCase := []string{
		"FB",
		"FoBa",
		"FoBaT",
	}
	expectList := [][]bool{
		[]bool{true, false, true, true, false},
		[]bool{true, false, true, false, false},
		[]bool{false, true, false, false, false},
	}

	for i, source := range sourceCase {
		expect := expectList[i]
		pattern := patternCase[i]

		target := camelMatch(source, pattern)
		for i := 0; i < len(expect); i++ {
			if target[i] != expect[i] {
				t.Error("Run Test_CamelMatch Fail", expect, target)
			}
		}
	}

	t.Log("Run Test_CamelMatch Success")
}
