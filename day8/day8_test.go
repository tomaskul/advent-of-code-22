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

func Test_GetVisibilityGrid_MatchesExpected(t *testing.T) {
	input := [][]int{
		0: {3, 0, 3, 7, 3},
		1: {2, 5, 5, 1, 2},
		2: {6, 5, 3, 3, 2},
		3: {3, 3, 5, 4, 9},
		4: {3, 5, 3, 9, 0}}
	expected := [][]int{
		0: {1, 1, 1, 1, 1},
		1: {1, 1, 1, 0, 1},
		2: {1, 1, 0, 1, 1},
		3: {1, 0, 1, 0, 1},
		4: {1, 1, 1, 1, 1},
	}

	visibilityGrid, totalVisible := getVisibilityGrid(input)
	if !reflect.DeepEqual(expected, visibilityGrid) {
		t.Errorf("Expected:\n%v\ngot:\n%v", expected, visibilityGrid)
	}
	if totalVisible != 21 {
		t.Errorf("Expected total visible: %d, got: %d", 21, totalVisible)
	}
}
