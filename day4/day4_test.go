package day4

import "testing"

func Test_NewCleanUpAssignment_MatchesExpected(t *testing.T) {
	testCases := map[string]*CleanUpAssignment{
		"1-5":    {LowerBound: 1, UpperBound: 5},
		"7-68":   {LowerBound: 7, UpperBound: 68},
		"10-891": {LowerBound: 10, UpperBound: 891},
		"45-58":  {LowerBound: 45, UpperBound: 58},
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual, err := newCleanUpAssignment(input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if actual == nil {
				t.Fatalf("Actual is nil")
			}

			if actual.LowerBound != expected.LowerBound {
				t.Errorf("Expected LB: %d, got: %d", expected.LowerBound, actual.LowerBound)
			}
			if actual.UpperBound != expected.UpperBound {
				t.Errorf("Expected UB: %d, got: %d", expected.UpperBound, actual.UpperBound)
			}
		})
	}
}

func Test_IsFullOverlap_IsTrue(t *testing.T) {
	testCases := map[string][]*CleanUpAssignment{
		"2-8,3-7":   {{LowerBound: 2, UpperBound: 8}, {LowerBound: 3, UpperBound: 7}},
		"6-6,4-6":   {{LowerBound: 6, UpperBound: 6}, {LowerBound: 4, UpperBound: 6}},
		"7-38,7-38": {{LowerBound: 7, UpperBound: 38}, {LowerBound: 7, UpperBound: 38}},
	}

	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			if !isFullOverlap(input[0], input[1]) {
				t.Errorf("Expected: true, got: false")
			}
		})
	}
}

func Test_IsFullOverlap_IsFalse(t *testing.T) {
	testCases := map[string][]*CleanUpAssignment{
		"2-4,6-8":     {{LowerBound: 2, UpperBound: 4}, {LowerBound: 6, UpperBound: 8}},
		"2-3,4-5":     {{LowerBound: 2, UpperBound: 3}, {LowerBound: 4, UpperBound: 5}},
		"5-43,3-9":    {{LowerBound: 5, UpperBound: 43}, {LowerBound: 3, UpperBound: 9}},
		"5-7,7-9":     {{LowerBound: 5, UpperBound: 7}, {LowerBound: 7, UpperBound: 9}},
		"2-6,4-8":     {{LowerBound: 2, UpperBound: 6}, {LowerBound: 4, UpperBound: 8}},
		"10-17,15-35": {{LowerBound: 10, UpperBound: 17}, {LowerBound: 15, UpperBound: 35}},
		"79-79,3-78":  {{LowerBound: 79, UpperBound: 79}, {LowerBound: 3, UpperBound: 78}},
		"42-53,13-41": {{LowerBound: 42, UpperBound: 53}, {LowerBound: 13, UpperBound: 41}},
	}

	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			if isFullOverlap(input[0], input[1]) {
				t.Errorf("Expected: false, got: true")
			}
		})
	}
}

func Test_IsPartialOverlap_IsTrue(t *testing.T) {
	testCases := map[string][]*CleanUpAssignment{
		"2-8,3-7":     {{LowerBound: 2, UpperBound: 8}, {LowerBound: 3, UpperBound: 7}},
		"6-6,4-6":     {{LowerBound: 6, UpperBound: 6}, {LowerBound: 4, UpperBound: 6}},
		"7-38,7-38":   {{LowerBound: 7, UpperBound: 38}, {LowerBound: 7, UpperBound: 38}},
		"5-43,3-9":    {{LowerBound: 5, UpperBound: 43}, {LowerBound: 3, UpperBound: 9}},
		"5-7,7-9":     {{LowerBound: 5, UpperBound: 7}, {LowerBound: 7, UpperBound: 9}},
		"2-6,4-8":     {{LowerBound: 2, UpperBound: 6}, {LowerBound: 4, UpperBound: 8}},
		"10-17,15-35": {{LowerBound: 10, UpperBound: 17}, {LowerBound: 15, UpperBound: 35}},
	}

	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			if !isPartialOverlap(input[0], input[1]) {
				t.Errorf("Expected: true, got: false")
			}
		})
	}
}

func Test_IsPartialOverlap_IsFalse(t *testing.T) {
	testCases := map[string][]*CleanUpAssignment{
		"2-4,6-8":     {{LowerBound: 2, UpperBound: 4}, {LowerBound: 6, UpperBound: 8}},
		"2-3,4-5":     {{LowerBound: 2, UpperBound: 3}, {LowerBound: 4, UpperBound: 5}},
		"79-79,3-78":  {{LowerBound: 79, UpperBound: 79}, {LowerBound: 3, UpperBound: 78}},
		"42-53,13-41": {{LowerBound: 42, UpperBound: 53}, {LowerBound: 13, UpperBound: 41}},
	}

	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			if isPartialOverlap(input[0], input[1]) {
				t.Errorf("Expected: false, got: true")
			}
		})
	}
}
