package util

import "testing"

const SessionCookie = ""

func Test_GetRows_NoEmptyRows(t *testing.T) {
	t.Skip("Specify sessionCookie")
	testCases := map[string]string{
		"Day 4": "https://adventofcode.com/2022/day/4/input",
		"Day 5": "https://adventofcode.com/2022/day/4/input",
	}

	for name, url := range testCases {
		t.Run(name, func(t *testing.T) {
			rows := GetRows(url, SessionCookie)
			for i, v := range rows {
				if v == "" {
					t.Errorf("Unexpected empty row at index: %d (length: %d)", i, len(rows))
				}
			}
		})
	}
}
