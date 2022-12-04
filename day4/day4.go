package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

type CleanUpAssignment struct {
	LowerBound int
	UpperBound int
}

func DayFourPt1(sessionCookie string) {
	pairs := util.GetRows("https://adventofcode.com/2022/day/4/input", sessionCookie)

	fmt.Println("Finding in how many assignment pairs does one range fully contain the other...")
	fullOverlaps := 0
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		assignments := strings.Split(pair, ",")
		elf1Assignment, err := newCleanUpAssignment(assignments[0])
		if err != nil {
			fmt.Printf("Error getting elf1 assignment: %v", err)
			continue
		}
		elf2Assignment, err := newCleanUpAssignment(assignments[1])
		if err != nil {
			fmt.Printf("Error getting elf2 assignment: %v", err)
			continue
		}

		if isFullOverlap(elf1Assignment, elf2Assignment) {
			fullOverlaps++
		}
	}

	fmt.Printf("Full overlaps: %d\n", fullOverlaps)
}

func newCleanUpAssignment(assignment string) (*CleanUpAssignment, error) {
	parts := strings.Split(assignment, "-")
	lb, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	ub, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &CleanUpAssignment{
		LowerBound: lb,
		UpperBound: ub,
	}, nil
}

func isFullOverlap(elf1Assignment, elf2Assignment *CleanUpAssignment) bool {
	minLb := int(math.Min(float64(elf1Assignment.LowerBound), float64(elf2Assignment.LowerBound)))
	if elf1Assignment.LowerBound >= minLb && elf2Assignment.LowerBound >= minLb {
		maxUb := int(math.Max(float64(elf1Assignment.UpperBound), float64(elf1Assignment.UpperBound)))
		if elf1Assignment.UpperBound <= maxUb && elf2Assignment.UpperBound <= maxUb {
			return true
		}
	}
	return false
}
