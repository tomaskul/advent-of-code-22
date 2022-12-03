package day2

import "testing"

var (
	drawTestCases = map[string][]string{
		"A X": {"A", "X"},
		"B Y": {"B", "Y"},
		"C Z": {"C", "Z"},
	}

	winTestCases = map[string][]string{
		"A Y": {"A", "Y"},
		"B Z": {"B", "Z"},
		"C X": {"C", "X"},
	}

	lossTestCases = map[string][]string{
		"A Z": {"A", "Z"},
		"B X": {"B", "X"},
		"C Y": {"C", "Y"},
	}
)

func Test_EvaluatePt1_ExpectedDraw(t *testing.T) {
	for name, tc := range drawTestCases {
		t.Run(name, func(t *testing.T) {
			opp, you, err := evaluatePt1(tc[0], tc[1])
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if opp != you {
				t.Errorf("Expected DRAW, got: %d v %d", opp, you)
			}
		})
	}
}

func Test_EvaluatePt1_ExpectedWin(t *testing.T) {
	for name, tc := range winTestCases {
		t.Run(name, func(t *testing.T) {
			opp, you, err := evaluatePt1(tc[0], tc[1])
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if you <= opp {
				t.Errorf("Expected WIN, got: %d v %d", opp, you)
			}
		})
	}
}

func Test_EvaluatePt1_ExpectedLoss(t *testing.T) {
	for name, tc := range lossTestCases {
		t.Run(name, func(t *testing.T) {
			opp, you, err := evaluatePt1(tc[0], tc[1])
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if you >= opp {
				t.Errorf("Expected LOSS, got: %d v %d", opp, you)
			}
		})
	}
}

type Case struct {
	OpponentHand          string
	YourHand              string
	OpponentExpectedScore int
	YourExpectedScore     int
}

func Test_EvaluatePt1_ExpectedScores(t *testing.T) {
	testCases := map[string]Case{
		"Rock Draw":     {OpponentHand: "A", YourHand: "X", OpponentExpectedScore: 4, YourExpectedScore: 4},
		"Paper Draw":    {OpponentHand: "B", YourHand: "Y", OpponentExpectedScore: 5, YourExpectedScore: 5},
		"Scissors Draw": {OpponentHand: "C", YourHand: "Z", OpponentExpectedScore: 6, YourExpectedScore: 6},

		"A Y": {OpponentHand: "A", YourHand: "Y", OpponentExpectedScore: 1, YourExpectedScore: 8},
		"A Z": {OpponentHand: "A", YourHand: "Z", OpponentExpectedScore: 7, YourExpectedScore: 3},

		"B X": {OpponentHand: "B", YourHand: "X", OpponentExpectedScore: 8, YourExpectedScore: 1},
		"B Z": {OpponentHand: "B", YourHand: "Z", OpponentExpectedScore: 2, YourExpectedScore: 9},

		"C X": {OpponentHand: "C", YourHand: "X", OpponentExpectedScore: 3, YourExpectedScore: 7},
		"C Y": {OpponentHand: "C", YourHand: "Y", OpponentExpectedScore: 9, YourExpectedScore: 2},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			opp, you, err := evaluatePt1(tc.OpponentHand, tc.YourHand)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if opp != tc.OpponentExpectedScore {
				t.Errorf("Expected: %d, got: %d", tc.OpponentExpectedScore, opp)
			}
			if you != tc.YourExpectedScore {
				t.Errorf("Expected: %d, got: %d", tc.YourExpectedScore, you)
			}
		})
	}
}
