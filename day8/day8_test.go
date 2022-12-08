package day8

import (
	"reflect"
	"testing"
)

func Test_GetIntRow_MatchesExpected(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	actual := getIntRow("1234567890")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

type GridTestCase struct {
	Name          string
	Input         [][]int
	Expected      [][]int
	ExpectedCount int
}

func Test_GetVisibilityGrid_MatchesExpected(t *testing.T) {
	testCases := []GridTestCase{
		{
			Name: "Example",
			Input: [][]int{
				0: {3, 0, 3, 7, 3},
				1: {2, 5, 5, 1, 2},
				2: {6, 5, 3, 3, 2},
				3: {3, 3, 5, 4, 9},
				4: {3, 5, 3, 9, 0},
			},
			Expected: [][]int{
				0: {1, 1, 1, 1, 1},
				1: {1, 1, 1, 0, 1},
				2: {1, 1, 0, 1, 1},
				3: {1, 0, 1, 0, 1},
				4: {1, 1, 1, 1, 1},
			},
			ExpectedCount: 21,
		},
		{
			Name: "Real data sample",
			Input: [][]int{
				0: {2, 0, 1, 2, 2, 1, 2, 2, 2, 3, 1},
				1: {2, 0, 1, 1, 1, 0, 2, 2, 3, 2, 1},
				2: {0, 2, 2, 0, 2, 0, 0, 2, 1, 3, 0},
				3: {2, 2, 0, 1, 0, 1, 3, 0, 0, 1, 0},
				4: {0, 1, 0, 0, 1, 3, 2, 0, 3, 2, 1},
				5: {2, 1, 2, 2, 2, 2, 2, 3, 2, 3, 1},
				6: {0, 0, 0, 2, 1, 0, 3, 3, 0, 2, 1},
			},
			Expected: [][]int{
				0: {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				1: {1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
				2: {1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
				3: {1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1},
				4: {1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1},
				5: {1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1},
				6: {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			ExpectedCount: 50,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			visibilityGrid, totalVisible := getVisibilityGrid(tC.Input)
			if !reflect.DeepEqual(tC.Expected, visibilityGrid) {
				t.Errorf("Expected:\n%v\ngot:\n%v", tC.Expected, visibilityGrid)
			}
			if totalVisible != tC.ExpectedCount {
				t.Errorf("Expected total visible: %d, got: %d", tC.ExpectedCount, totalVisible)
			}
		})
	}
}
