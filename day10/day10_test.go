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

func Test_RunCycles(t *testing.T) {
	testCases := []struct {
		desc          string
		instructions  []*Instruction
		signalAtCycle []int
		signalValue   []int
	}{
		{
			desc:          "Small example",
			instructions:  []*Instruction{{Cycles: 1, Value: 0}, {Cycles: 2, Value: 3}, {Cycles: 2, Value: -5}},
			signalAtCycle: []int{1, 2, 3, 4, 5, 6},
			signalValue:   []int{1, 2, 3, 16, 20, -6}, // values: []int{1, 1, 1, 4, 4, -1}
		},
		{
			desc: "2nd example",
			instructions: getInstructionsByValue([]int{15, -11, 6, -3, 5, -1, 8, 13, 4, 0, -1, 5, -1,
				5, -1, 5, -1, 5, -1, -35, 1, 24, -19, 1, 16, -11, 0, 0, 21, -15, 0, 0, -3, 9, 1, -3, 8, 1, 5,
				0, 0, 0, 0, 0, -36, 0, 1, 7, 0, 0, 0, 2, 6, 0, 0, 0, 0, 0, 1, 0, 0, 7, 1, 0, -13, 13, 7, 0, 1,
				-33, 0, 0, 0, 2, 0, 0, 0, 8, 0, -1, 2, 1, 0, 17, -9, 1, 1, -3, 11, 0, 0, 1, 0, 1, 0, 0, -13,
				-19, 1, 3, 26, -30, 12, -1, 3, 1, 0, 0, 0, -9, 18, 1, 2, 0, 0, 9, 0, 0, 0, -1, 2, -37, 1, 3, 0,
				15, -21, 22, -6, 1, 0, 2, 1, 0, -10, 0, 0, 20, 1, 2, 2, -6, -11, 0, 0, 0}),
			signalAtCycle: []int{20, 60, 100, 140, 180, 220},
			signalValue:   []int{420, 1140, 1800, 2940, 2880, 3960},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer Reset()
			actual := runCycles(tC.instructions, tC.signalAtCycle)
			if len(actual) != len(tC.signalAtCycle) {
				t.Fatalf("Expected: %d, got: %d", len(tC.signalAtCycle), len(actual))
			}
			if !reflect.DeepEqual(actual, tC.signalValue) {
				t.Errorf("Expected:\n%v\ngot:\n%v", tC.signalValue, actual)
			}
		})
	}
}

func getInstructionsByValue(values []int) []*Instruction {
	result := make([]*Instruction, len(values))
	for i := 0; i < len(values); i++ {
		var instruction *Instruction
		if values[i] == 0 {
			instruction = &Instruction{Cycles: 1, Value: 0}
		} else {
			instruction = &Instruction{Cycles: 2, Value: values[i]}
		}
		result[i] = instruction
	}
	return result
}
