package day2

import (
	"fmt"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

const (
	DrawScore = 3
	WinScore  = 6
)

var yourMatrix = [][]string{
	0: {"Y", "Z", "X"},
	1: {"X", "Y", "Z"},
	2: {"Z", "X", "Y"},
}

var charLookup = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"X": 0,
	"Y": 1,
	"Z": 2,
}

func Solution(sessionCookie, pt1Text, pt2Text string) {
	bytes, _ := util.GetData("https://adventofcode.com/2022/day/2/input", sessionCookie)
	rounds := strings.Split(string(bytes), "\n")

	fmt.Printf(pt1Text)
	solve(rounds, evaluatePt1)
	fmt.Printf(pt2Text)
	solve(rounds, evaluatePt2)
}

func solve(rounds []string, evalFunc func(string, string) (int, int, error)) {
	opponentScore, myScore := 0, 0
	for _, round := range rounds {
		parts := strings.Split(round, " ")
		if len(parts) < 2 {
			continue
		}
		opp, you, _ := evalFunc(parts[0], parts[1])
		opponentScore += opp
		myScore += you
	}

	printSummary(opponentScore, myScore)
}

func printSummary(opponentScore, myScore int) {
	message := ""
	if opponentScore > myScore {
		message = "You LOST!"
	} else {
		message = "You WON!"
	}

	fmt.Printf("%s\nOpponent      You\n%d    vs   %d\n", message, opponentScore, myScore)
}

func evaluatePt1(opponentHand, yourHand string) (int, int, error) {
	opp, _ := getHandScore(opponentHand)
	you, _ := getHandScore(yourHand)
	oppIdx := charLookup[opponentHand]
	youIdx := charLookup[yourHand]

	outcome := yourMatrix[oppIdx][youIdx]
	switch outcome {
	case "Y":
		opp += DrawScore
		you += DrawScore
	case "Z":
		you += WinScore
	case "X":
		opp += WinScore
	}
	return opp, you, nil
}

func getHandScore(input string) (int, error) {
	switch input {
	case "A", "X": // Rock
		return 1, nil
	case "B", "Y": // Paper
		return 2, nil
	case "C", "Z": // Scissors
		return 3, nil
	default:
		return -1, fmt.Errorf("unknown input '%s'", input)
	}
}

func evaluatePt2(opponentHand, targetOutcome string) (int, int, error) {
	opp, _ := getHandScore(opponentHand)
	options := yourMatrix[charLookup[opponentHand]]

	var yourHandIdx int
	for idx, v := range options {
		if v == targetOutcome {
			yourHandIdx = idx
		}
	}

	you := yourHandIdx + 1
	switch targetOutcome {
	case "X":
		// Lose.
		opp += WinScore
	case "Y":
		// Draw.
		opp += DrawScore
		you += DrawScore
	case "Z":
		// Win.
		you += WinScore
	}
	return opp, you, nil
}
