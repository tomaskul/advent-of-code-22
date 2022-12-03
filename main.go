package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomaskul/advent-of-code-22/day1"
)

const DaysSolved = 1

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
		fmt.Printf("\t=== Part 1 ===\n")
		day1.DayOnePt1(sessionCookie)
		fmt.Printf("\n\t=== Part 2 ===\n")
		day1.DayOnePt2(sessionCookie)
	}
}
