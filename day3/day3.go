package day3

import (
	"fmt"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

func DayThreePt1(sessionCookie string) {
	dataBytes, _ := util.GetData("https://adventofcode.com/2022/day/3/input", sessionCookie)
	rucksacks := strings.Split(string(dataBytes), "\n")

	for index, rucksack := range rucksacks {
		compartments, err := splitIntoCompartments(rucksack)
		if err != nil {
			fmt.Printf("rucksacks[%d]: %v", index, err)
			continue
		}

		_, err = findMatch(compartments[0], compartments[1])
		if err != nil {
			fmt.Printf("rucksacks[%d]: %v", index, err)
			continue
		}
	}
}

func splitIntoCompartments(rucksack string) ([]string, error) {
	if len(rucksack)%2 != 0 {
		return nil, fmt.Errorf("rucksack: '%s' not divisible into compartments", rucksack)
	}
	midPoint := len(rucksack) / 2
	return []string{rucksack[:midPoint], rucksack[midPoint:]}, nil
}

func findMatch(compartment1, compartment2 string) (rune, error) {
	for i := 0; i < len(compartment1); i++ {
		for j := 0; j < len(compartment2); j++ {
			if compartment2[j] == compartment1[i] {
				return rune(compartment1[i]), nil
			}
		}
	}

	return rune(' '), fmt.Errorf("no matches found")
}

/*
func getItemPriority(item rune) (int, error) {
	isLowerCase, err := regexp.MatchString("[a-z]", string(item))
	if err != nil {
		return -1, err
	}
	if isLowerCase {

	} else {

	}
}
*/
