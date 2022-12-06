package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomaskul/advent-of-code-22/day1"
	"github.com/tomaskul/advent-of-code-22/day2"
	"github.com/tomaskul/advent-of-code-22/day3"
	"github.com/tomaskul/advent-of-code-22/day4"
	"github.com/tomaskul/advent-of-code-22/day5"
	"github.com/tomaskul/advent-of-code-22/day6"
)

const (
	DaysSolved = 6

	Part1Text = "\t=== Part 1 ===\n"
	Part2Text = "\n\t=== Part 2 ===\n"
)

func main() {
	var sessionCookie string
	var day int
	flag.StringVar(&sessionCookie, "s", "", "Session cookie to auth & retrieve user specific problem.")
	flag.IntVar(&day, "day", 1, "Day number for which to run solution for.")
	flag.Parse()

	if sessionCookie == "" {
		fmt.Println("Invalid session cookie supplied")
		os.Exit(-1)
	}
	if day < 1 || day > DaysSolved {
		fmt.Printf("Invalid day supplied: %d (range: [1:%d]\n", day, DaysSolved)
		os.Exit(-1)
	}

	switch day {
	case 1:
		day1.Solution(sessionCookie, Part1Text, Part2Text)
	case 2:
		day2.Solution(sessionCookie, Part1Text, Part2Text)
	case 3:
		day3.Solution(sessionCookie, Part1Text, Part2Text)
	case 4:
		day4.Solution(sessionCookie, Part1Text, Part2Text)
	case 5:
		day5.Solution(sessionCookie, Part1Text, Part2Text)
	case 6:
		day6.Solution(sessionCookie, Part1Text, Part2Text)
	}
}
