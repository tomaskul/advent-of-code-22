package day10

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_ParseInts(t *testing.T) {
	testCases := []struct {
		desc  string
		value int
	}{
		{desc: "0", value: 0},
		{desc: "8", value: 8},
		{desc: "-1", value: -1},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := strconv.Atoi(tC.desc)
			if err != nil {
				t.Fatalf("Error when parsing: '%s': %v", tC.desc, err)
			}
			if actual != tC.value {
				t.Errorf("Expected: %d, got: %d", tC.value, actual)
			}
		})
	}
}

var example_1 = []string{"noop", "addx 3", "addx -5"}
var example_2 = []string{"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8", "addx 13", "addx 4",
	"noop", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx -35", "addx 1",
	"addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop", "noop", "addx 21", "addx -15", "noop", "noop", "addx -3",
	"addx 9", "addx 1", "addx -3", "addx 8", "addx 1", "addx 5", "noop", "noop", "noop", "noop", "noop", "addx -36", "noop",
	"addx 1", "addx 7", "noop", "noop", "noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop", "noop", "addx 1", "noop",
	"noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7", "noop", "addx 1", "addx -33", "noop", "noop", "noop",
	"addx 2", "noop", "noop", "noop", "addx 8", "noop", "addx -1", "addx 2", "addx 1", "noop", "addx 17", "addx -9", "addx 1",
	"addx 1", "addx -3", "addx 11", "noop", "noop", "addx 1", "noop", "addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1",
	"addx 3", "addx 26", "addx -30", "addx 12", "addx -1", "addx 3", "addx 1", "noop", "noop", "noop", "addx -9", "addx 18",
	"addx 1", "addx 2", "noop", "noop", "addx 9", "noop", "noop", "noop", "addx -1", "addx 2", "addx -37", "addx 1", "addx 3",
	"noop", "addx 15", "addx -21", "addx 22", "addx -6", "addx 1", "noop", "addx 2", "addx 1", "noop", "addx -10", "noop", "noop",
	"addx 20", "addx 1", "addx 2", "addx 2", "addx -6", "addx -11", "noop", "noop", "noop"}
var puzzleInput = []string{"addx 1", "addx 4", "addx 21", "addx -20", "addx 4", "noop", "noop", "addx 5", "addx 3", "noop", "addx 2",
	"addx 1", "noop", "noop", "addx 4", "noop", "noop", "noop", "addx 3", "addx 5", "addx 2", "addx 1", "noop", "addx -37", "addx 22",
	"addx -4", "addx -14", "addx 2", "addx 5", "addx 3", "addx -2", "addx 2", "addx 5", "addx 2", "addx -15", "addx 32", "addx -14", "addx 5",
	"addx 2", "addx 3", "noop", "addx -13", "addx -2", "addx 18", "addx -36", "noop", "addx 11", "addx -7", "noop", "noop", "addx 6", "addx 22",
	"addx -21", "addx 3", "addx 2", "addx 4", "noop", "noop", "noop", "addx 5", "addx -16", "addx 17", "addx 2", "addx 5", "addx -11", "addx 15",
	"addx -15", "addx -24", "noop", "noop", "addx 7", "addx 2", "addx -6", "addx 9", "noop", "addx 5", "noop", "addx -3", "addx 4", "addx 2",
	"noop", "noop", "addx 7", "noop", "noop", "noop", "addx 5", "addx -28", "addx 29", "noop", "addx 3", "addx -7", "addx -29", "noop", "addx 7",
	"addx -2", "addx 2", "addx 5", "addx 2", "addx -3", "addx 4", "addx 5", "addx 2", "addx 8", "addx -30", "addx 25", "addx 7", "noop", "noop",
	"addx 3", "addx -2", "addx 2", "addx -10", "addx -24", "addx 2", "noop", "noop", "addx 2", "noop", "addx 3", "addx 2", "noop", "addx 3",
	"addx 2", "addx 5", "addx 2", "noop", "addx 1", "noop", "addx 2", "addx 8", "noop", "noop", "addx -1", "addx -9", "addx 14", "noop",
	"addx 1", "noop", "noop"}

func Test_RunCycles(t *testing.T) {
	testCases := []struct {
		desc          string
		instructions  []string
		signalAtCycle []int
		signalValue   []int
	}{
		{
			desc:          "Small example",
			instructions:  example_1,
			signalAtCycle: []int{1, 2, 3, 4, 5, 6},
			signalValue:   []int{1, 2, 3, 16, 20, -6},
		},
		{
			desc:          "2nd example",
			instructions:  example_2,
			signalAtCycle: []int{20, 60, 100, 140, 180, 220},
			signalValue:   []int{420, 1140, 1800, 2940, 2880, 3960},
		},
		{
			desc:          "puzzle input",
			instructions:  puzzleInput,
			signalAtCycle: []int{20, 60, 100, 140, 180, 220},
			signalValue:   []int{420, 1260, 2100, 2380, 3780, 4620},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer Reset()

			actual := runCycles(parseInstructions(tC.instructions), tC.signalAtCycle, nil)
			if len(actual) != len(tC.signalAtCycle) {
				t.Fatalf("Expected: %d, got: %d", len(tC.signalAtCycle), len(actual))
			}
			if !reflect.DeepEqual(actual, tC.signalValue) {
				t.Errorf("Expected:\n%v\ngot:\n%v", tC.signalValue, actual)
			}
		})
	}
}

func Test_RunCyles_Render(t *testing.T) {
	testCases := []struct {
		desc          string
		instructions  []string
		signalAtCycle []int
	}{
		{
			desc:          "example 2",
			instructions:  example_2,
			signalAtCycle: []int{40, 80, 120, 160, 200, 240},
		},
		/*{
			desc:          "puzzle input",
			instructions:  puzzleInput,
			signalAtCycle: []int{40, 80, 120, 160, 200, 240},
		},*/
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer Reset()

			actual := runCycles(parseInstructions(tC.instructions), tC.signalAtCycle, renderSignalToConsole)
			if len(actual) != len(tC.signalAtCycle) {
				t.Fatalf("Expected: %d, got: %d", len(tC.signalAtCycle), len(actual))
			}
		})
	}
}
