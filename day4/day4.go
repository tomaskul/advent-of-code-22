package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

type CleanUpAssignment struct {
	LowerBound int
	UpperBound int
}

func Solution(sessionCookie, pt1Text, pt2Text string) {
	pairs := util.GetRows("https://adventofcode.com/2022/day/4/input", sessionCookie)

	fmt.Printf(pt1Text)
	fmt.Printf("Full overlaps: %d\n", solve(pairs, isFullOverlap))
	fmt.Printf(pt2Text)
	fmt.Printf("Partial overlaps: %d\n", solve(pairs, isPartialOverlap))
}

func solve(pairs []string, evalFunc func(*CleanUpAssignment, *CleanUpAssignment) bool) int {
	count := 0
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		assignments := strings.Split(pair, ",")
		elf1Assignment, _ := newCleanUpAssignment(assignments[0])
		elf2Assignment, _ := newCleanUpAssignment(assignments[1])

		if evalFunc(elf1Assignment, elf2Assignment) {
			count++
		}
	}

	return count
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
	if (elf1Assignment.LowerBound <= elf2Assignment.LowerBound && elf1Assignment.UpperBound >= elf2Assignment.UpperBound) ||
		(elf2Assignment.LowerBound <= elf1Assignment.LowerBound && elf2Assignment.UpperBound >= elf1Assignment.UpperBound) {
		return true
	}

	return false
}

func isPartialOverlap(elf1Assignment, elf2Assignment *CleanUpAssignment) bool {
	if isFullOverlap(elf1Assignment, elf2Assignment) {
		return true
	}

	if (elf1Assignment.LowerBound >= elf2Assignment.LowerBound && elf1Assignment.UpperBound >= elf2Assignment.UpperBound) ||
		(elf2Assignment.LowerBound >= elf1Assignment.LowerBound && elf2Assignment.UpperBound >= elf1Assignment.UpperBound) {
		return true
	}
	return false
}
