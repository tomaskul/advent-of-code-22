package day6

import (
	"testing"
)

func Test_FindStartOfPacketMarker_MatchesExpected(t *testing.T) {
	testCases := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual := findStartOfPacketMarker([]byte(input))
			if actual != expected {
				t.Errorf("Expected: %d, got: %d", expected, actual)
			}
		})
	}
}
