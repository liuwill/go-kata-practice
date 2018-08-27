package weekly_contest_93

import (
	"testing"
)

func Test_MinRefuelStops(t *testing.T) {
	destination := 100
	startFuel := 10
	stations := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	target := minRefuelStops(destination, startFuel, stations)
	expect := 2

	if expect != target {
		t.Error("Translate minRefuelStops Fail", target)
	}
	t.Log("Translate Test_MinRefuelStops Success")
}

func Test_MinRefuelStopsEmpty(t *testing.T) {
	destination := 1
	startFuel := 1
	stations := [][]int{}
	target := minRefuelStops(destination, startFuel, stations)
	expect := 0

	if expect != target {
		t.Error("Translate Test_MinRefuelStopsEmpty Fail", target)
	}
	t.Log("Translate Test_MinRefuelStopsEmpty Success")
}

func Test_MinRefuelStopsFail(t *testing.T) {
	destination := 20
	startFuel := 5
	stations := [][]int{{10, 20}}
	target := minRefuelStops(destination, startFuel, stations)
	expect := -1

	if expect != target {
		t.Error("Translate Test_MinRefuelStopsFail Fail", target)
	}
	t.Log("Translate Test_MinRefuelStopsFail Success")
}

func Test_MinRefuelStopsExtram(t *testing.T) {
	destination := 1000
	startFuel := 10
	stations := [][]int{
		{7, 217}, {10, 211}, {17, 146}, {21, 232}, {25, 310}, {48, 175},
		{53, 23}, {63, 158}, {71, 292}, {96, 85}, {100, 302}, {102, 295},
		{116, 110}, {122, 46}, {131, 20}, {132, 19}, {141, 13}, {163, 85},
		{169, 263}, {179, 233}, {191, 169}, {215, 163}, {224, 231}, {228, 282},
		{256, 115}, {259, 3}, {266, 245}, {283, 331}, {299, 21}, {310, 224},
		{315, 188}, {328, 147}, {345, 74}, {350, 49}, {379, 79}, {387, 276},
		{391, 92}, {405, 174}, {428, 307}, {446, 205}, {448, 226}, {452, 275},
		{475, 325}, {492, 310}, {496, 94}, {499, 313}, {500, 315}, {511, 137},
		{515, 180}, {519, 6}, {533, 206}, {536, 262}, {553, 326}, {561, 103},
		{564, 115}, {582, 161}, {593, 236}, {599, 216}, {611, 141}, {625, 137},
		{626, 231}, {628, 66}, {646, 197}, {665, 103}, {668, 8}, {691, 329},
		{699, 246}, {703, 94}, {724, 277}, {729, 75}, {735, 23}, {740, 228},
		{761, 73}, {770, 120}, {773, 82}, {774, 297}, {780, 184}, {791, 239},
		{801, 85}, {805, 156}, {837, 157}, {844, 259}, {849, 2}, {860, 115},
		{874, 311}, {877, 172}, {881, 109}, {888, 321}, {894, 302}, {899, 321},
		{908, 76}, {916, 241}, {924, 301}, {933, 56}, {960, 29}, {979, 319},
		{983, 325}, {988, 190}, {995, 299}, {996, 97},
	}
	target := minRefuelStops(destination, startFuel, stations)
	expect := 4

	if expect != target {
		t.Error("Translate Test_MinRefuelStopsExtram Fail", target)
	}
	t.Log("Translate Test_MinRefuelStopsExtram Success")
}
