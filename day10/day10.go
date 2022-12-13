package day10

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tomaskul/advent-of-code-22/util"
)

type Instruction struct {
	Value  int
	Cycles int
}

var registerX = 1

func Solution(sessionCookie, pt1Text, pt2Text string) {
	input := util.GetRows("https://adventofcode.com/2022/day/10/input", sessionCookie)
	instructions := parseInstructions(input)
	output := runCycles(instructions, []int{20, 60, 100, 140, 180, 220}, nil)
	sum := 0
	for _, v := range output {
		sum += v
	}
	fmt.Printf(pt1Text)
	fmt.Printf("Sum of 20, 60, 100, 140, 180, 220 signal strengths: %d\n", sum)

	// TODO: Work in progress.
	//fmt.Printf(pt2Text)
	//runCycles(instructions, []int{40, 80, 120, 160, 200, 240}, renderSignalToConsole)
}

func parseInstructions(instructions []string) []*Instruction {
	result := make([]*Instruction, len(instructions))

	for index, instruction := range instructions {
		parts := strings.Split(instruction, " ")

		instructionCycles := 1
		value := 0
		if parts[0] == "addx" {
			instructionCycles = 2
			intValue, err := strconv.Atoi(parts[1])
			value = intValue
			if err != nil {
				fmt.Printf("Couldn't convert '%s' to int...", parts[1])
				continue
			}
		}

		result[index] = &Instruction{Cycles: instructionCycles, Value: value}
	}

	return result
}

func runCycles(instructions []*Instruction, signalAtCycle []int, renderFunc func(int, int)) []int {
	signalAtValues := []int{}
	cycleCounter := 0
	timeout := time.Now().Add(time.Minute * 5)
	for {
		cycleCounter++
		if len(signalAtValues) == len(signalAtCycle) || time.Now().After(timeout) {
			break
		}

		if instructions[0].Cycles == 0 {
			registerX += instructions[0].Value
			instructions = instructions[1:]
		}
		if len(instructions) > 0 {
			instructions[0].Cycles = instructions[0].Cycles - 1
		}

		// Log - not part of cycle logic.
		if signalAtCycle[0] == cycleCounter {
			signalAtValues = append(signalAtValues, registerX*cycleCounter)
			signalAtCycle = append(signalAtCycle[1:], signalAtCycle[0])
		}

		if renderFunc != nil {
			renderFunc(registerX, cycleCounter)
		}
	}

	return signalAtValues
}

func renderSignalToConsole(xPos, cycleCounter int) {
	screenWidth := 40
	if isPixelLit(xPos, cycleCounter, screenWidth) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if isNewLine(cycleCounter, screenWidth) {
		fmt.Println()
	}
}

func isNewLine(cycleCounter, screenWidth int) bool {
	return cycleCounter == screenWidth || cycleCounter == screenWidth*2 || cycleCounter == screenWidth*3 ||
		cycleCounter == screenWidth*4 || cycleCounter == screenWidth*5 || cycleCounter == screenWidth*6
}

func isPixelLit(xPos, cycleCounter, screenWidth int) bool {
	multiples := cycleCounter / screenWidth

	adjustedCycleCounter := (cycleCounter - (screenWidth * multiples))

	return adjustedCycleCounter == xPos-1 || adjustedCycleCounter == xPos || adjustedCycleCounter == xPos+1
}

// Resets day 10 solution state.
func Reset() {
	registerX = 1
}
