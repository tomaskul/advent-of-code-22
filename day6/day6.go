package day6

import (
	"fmt"

	"github.com/tomaskul/advent-of-code-22/util"
)

func Solution(sessionCookie, pt1Text, pt2Text string) {
	fmt.Printf(pt1Text)
	data, _ := util.GetData("https://adventofcode.com/2022/day/6/input", sessionCookie)
	fmt.Printf("How many characters need to be processed before the first start-of-packet marker is detected?: %d\n", findStartOfPacketMarker(data))

	fmt.Printf(pt2Text)
	fmt.Printf("How many characters need to be processed before the first start-of-message marker is detected?: %d\n", findStartOfMessageMarker(data))
}

func findStartOfPacketMarker(data []byte) int {
	return findMarker(data, 4)
}

func findMarker(data []byte, markerLength int) int {
	marker := []byte{}
	for i := 0; i < len(data); i++ {
		hasChar, furthestIndex := hasCharacter(marker, data[i])
		if hasChar {
			marker = marker[furthestIndex+1:]
		}

		marker = append(marker, data[i])

		if len(marker) == markerLength {
			return i + 1
		}
	}

	return -1
}

func hasCharacter(space []byte, character byte) (bool, int) {
	result, furthestIndex := false, 0
	for i := 0; i < len(space); i++ {
		if space[i] == character {
			result = true
			furthestIndex = i
		}
	}

	return result, furthestIndex
}

func findStartOfMessageMarker(data []byte) int {
	return findMarker(data, 14)
}
