package day3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

func Solution(sessionCookie, pt1Text, pt2Text string) {
	dataBytes, _ := util.GetData("https://adventofcode.com/2022/day/3/input", sessionCookie)
	rucksacks := strings.Split(string(dataBytes), "\n")

	fmt.Printf(pt1Text)
	dayThreePt1(rucksacks)
	fmt.Printf(pt2Text)
	dayThreePt2(rucksacks)
}

func dayThreePt1(rucksacks []string) {
	fmt.Println("Finding the items type that appears in both compartments of each rucksack...")
	sum := 0
	for index, rucksack := range rucksacks {
		if rucksack == "" {
			continue
		}
		compartments, err := splitIntoCompartments(rucksack)
		if err != nil {
			fmt.Printf("rucksacks[%d]: %v", index, err)
			continue
		}

		item, err := findMatch(compartments[0], compartments[1], "")
		if err != nil {
			fmt.Printf("rucksacks[%d]: %v", index, err)
			continue
		}

		priority, err := getItemPriority(item)
		if err != nil {
			fmt.Printf("rucksacks[%d]: %v", index, err)
			continue
		}

		sum += priority
	}

	fmt.Printf("Sum of the priorities of those item types: %d\n", sum)
}

func dayThreePt2(rucksacks []string) {
	fmt.Println("Finding the item types that corresponds to the badges of each three-Elf group...")
	groups := getGroupRucksacks(rucksacks)
	sum := 0
	for _, group := range groups {
		match, err := findMatch(group[0], group[1], group[2])
		if err != nil {
			fmt.Printf("Error finding match in group: %v", err)
			continue
		}

		priority, _ := getItemPriority(match)
		sum += priority
	}

	fmt.Printf("What is the sum of the priorities of those item types: %d\n", sum)
}

func splitIntoCompartments(rucksack string) ([]string, error) {
	if len(rucksack)%2 != 0 {
		return nil, fmt.Errorf("rucksack: '%s' not divisible into compartments", rucksack)
	}
	midPoint := len(rucksack) / 2
	return []string{rucksack[:midPoint], rucksack[midPoint:]}, nil
}

func findMatch(compartment1, compartment2, compartment3 string) (rune, error) {
	for i := 0; i < len(compartment1); i++ {
		if lookFor(compartment1[i], compartment2) {
			if compartment3 == "" {
				return rune(compartment1[i]), nil
			}
			if lookFor(compartment1[i], compartment3) {
				return rune(compartment1[i]), nil
			}
		}
	}

	return rune(' '), fmt.Errorf("no matches found")
}

func lookFor(target byte, searchSpace string) bool {
	for i := 0; i < len(searchSpace); i++ {
		if searchSpace[i] == target {
			return true
		}
	}
	return false
}

func getItemPriority(item rune) (int, error) {
	isLowerCase, _ := regexp.MatchString("[a-z]", string(item))
	if isLowerCase {
		return int(item) - 96, nil
	} else {
		return int(item) - 38, nil
	}
}

func getGroupRucksacks(rucksacks []string) map[int][]string {
	result := make(map[int][]string)
	groupIndex := 0
	groupMember := 0
	for _, rucksack := range rucksacks {
		if rucksack == "" {
			continue
		}
		if groupMember == 3 {
			groupIndex++
			groupMember = 0
		}

		groupMember++
		result[groupIndex] = append(result[groupIndex], rucksack)
	}

	return result
}
