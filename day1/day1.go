package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

const inputUrl = "https://adventofcode.com/2022/day/1/input"

func DayOnePt1(sessionCookie string) {
	data, _ := util.GetData(inputUrl, sessionCookie)
	totals, _ := parseData(data)

	fmt.Println("Finding the Elf carrying the most Calories...")

	highest, _ := findHighest(totals)
	fmt.Printf("Elf carrying most Calories, is carrying: %d calories.\n", highest)
}

func DayOnePt2(sessionCookie string) {
	data, _ := util.GetData(inputUrl, sessionCookie)
	totals, _ := parseData(data)

	fmt.Printf("Find the top three Elves carrying the most Calories...\n=== Top 3 ===\n")

	top1, top1Index := findHighest(totals)
	fmt.Printf("\t#1: %d\n", top1)

	top2, top2index := findHighest(append(totals[:top1Index], totals[top1Index+1:]...))
	fmt.Printf("\t#2: %d\n", top2)

	top3, _ := findHighest(append(totals[:top2index], totals[top2index+1:]...))
	fmt.Printf("\t#3: %d\n", top3)

	fmt.Printf("Total for top #3: %d\n", top1+top2+top3)
}

func parseData(input []byte) ([]int, error) {
	lines := strings.Split(string(input), "\n")

	result := make([]int, 0)
	currentTotal := 0
	for _, v := range lines {
		if v == "" {
			result = append(result, currentTotal)
			currentTotal = 0
			continue
		} else {
			value, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error parsing data!")
				return nil, err
			}
			currentTotal += value
		}
	}

	return result, nil
}

func findHighest(input []int) (int, int) {
	result, resultIndex := 0, 0

	for index, v := range input {
		if v > result {
			result = v
			resultIndex = index
		}
	}

	return result, resultIndex
}
