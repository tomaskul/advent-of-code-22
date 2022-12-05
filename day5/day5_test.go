package day5

import (
	"reflect"
	"testing"
)

func Test_NewInstruction_MatchesExpected(t *testing.T) {
	testCases := map[string]*Instruction{
		"move 3 from 2 to 1":  {Move: 3, From: 2, To: 1},
		"move 10 from 5 to 4": {Move: 10, From: 5, To: 4},
	}
	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual := newInstruction(input)
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Actual doesn't match expected")
			}
		})
	}
}
