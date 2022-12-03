package day3

import (
	"testing"
)

type Expected struct {
	Compartments  []string
	Match         rune
	MatchPriority int
}

var testInputs = map[string]Expected{
	"vJrwpWtwJgWrhcsFMMfFFhFp":         {Compartments: []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, Match: 'p', MatchPriority: 16},
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL": {Compartments: []string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, Match: 'L', MatchPriority: 38},
	"PmmdzqPrVvPwwTWBwg":               {Compartments: []string{"PmmdzqPrV", "vPwwTWBwg"}, Match: 'P', MatchPriority: 42},
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn":   {Compartments: []string{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"}, Match: 'v', MatchPriority: 22},
	"ttgJtRGJQctTZtZT":                 {Compartments: []string{"ttgJtRGJ", "QctTZtZT"}, Match: 't', MatchPriority: 20},
	"CrZsJsPPZsGzwwsLwLmpwMDw":         {Compartments: []string{"CrZsJsPPZsGz", "wwsLwLmpwMDw"}, Match: 's', MatchPriority: 19},
}

func Test_SplitIntoCompartments_MatchesExpected(t *testing.T) {
	for input, expected := range testInputs {
		t.Run(input, func(t *testing.T) {
			actual, err := splitIntoCompartments(input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if len(actual) != len(expected.Compartments) {
				t.Fatalf("Expected length: %d, got: %d", len(expected.Compartments), len(actual))
			}
			for i, compartment := range expected.Compartments {
				if actual[i] != compartment {
					t.Errorf("Expected[%d]: '%s', got: '%s'", i, compartment, actual[i])
				}
			}
		})
	}
}

func Test_FindMatch_MatchesExpected(t *testing.T) {
	for input, expected := range testInputs {
		t.Run(input, func(t *testing.T) {
			compartments, err := splitIntoCompartments(input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			actual, err := findMatch(compartments[0], compartments[1], "")
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if expected.Match != actual {
				t.Errorf("Expected: '%q', got: '%q'", expected.Match, actual)
			}
		})
	}
}

func Test_FindMatchThreeInputs_MatchExpected(t *testing.T) {
	testCases := map[rune][]string{
		'r': {"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
		'Z': {"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
	}

	for expected, input := range testCases {
		t.Run(string(expected), func(t *testing.T) {
			actual, err := findMatch(input[0], input[1], input[2])
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if expected != actual {
				t.Errorf("Expected: %q, got: %q", expected, actual)
			}
		})
	}
}

func Test_GetItemPriority_MatchesExpected(t *testing.T) {
	for input, expected := range testInputs {
		t.Run(input, func(t *testing.T) {
			compartments, err := splitIntoCompartments(input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			match, err := findMatch(compartments[0], compartments[1], "")
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			actual, err := getItemPriority(match)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if expected.MatchPriority != actual {
				t.Errorf("Expected: %d, got: %d", expected.MatchPriority, actual)
			}
		})
	}
}

func Test_GetGroupRucksacks_MatchesExpected(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	actual := getGroupRucksacks(input)

	if len(actual) != 2 {
		t.Errorf("Expected 2 groups, got: %d", len(actual))
	}
}
