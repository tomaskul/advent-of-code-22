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
		"2-8,3-7": {{LowerBound: 2, UpperBound: 8}, {LowerBound: 3, UpperBound: 7}},
		"6-6,4-6": {{LowerBound: 6, UpperBound: 6}, {LowerBound: 4, UpperBound: 6}},
	}

	for _, input := range testCases {
		if !isFullOverlap(input[0], input[1]) {
			t.Errorf("Expected: true, got: false")
		}
	}
}

func Test_IsFullOverlap_IsFalse(t *testing.T) {
	testCases := map[string][]*CleanUpAssignment{
		"2-4,6-8": {{LowerBound: 2, UpperBound: 4}, {LowerBound: 6, UpperBound: 8}},
		"2-3,4-5": {{LowerBound: 2, UpperBound: 3}, {LowerBound: 4, UpperBound: 5}},
		"5-7,7-9": {{LowerBound: 5, UpperBound: 7}, {LowerBound: 7, UpperBound: 9}},
		"2-6,4-8": {{LowerBound: 2, UpperBound: 6}, {LowerBound: 4, UpperBound: 8}},
	}

	for _, input := range testCases {
		if isFullOverlap(input[0], input[1]) {
			t.Errorf("Expected: false, got: true")
		}
	}
}
