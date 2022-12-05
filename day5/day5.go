package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

type Instruction struct {
	Move int
	From int
	To   int
}

func Solution(sessionCookie, pt1Text, _ string) {
	fmt.Printf(pt1Text)
	setInitialState()
	instructions := getInstructions(sessionCookie)

	fmt.Printf("%v\n", instructions[0])
	fmt.Printf("%v\n", instructions[len(instructions)-1])
}

func setInitialState() map[int]*util.Stack {
	state := map[int]*util.Stack{
		1: {}, 2: {}, 3: {},
		4: {}, 5: {}, 6: {},
		7: {}, 8: {}, 9: {},
	}
	state[1].PushMultiple([]string{"M", "J", "C", "B", "F", "R", "L", "H"})
	state[2].PushMultiple([]string{"Z", "C", "D"})
	state[3].PushMultiple([]string{"H", "J", "F", "C", "N", "G", "W"})
	state[4].PushMultiple([]string{"P", "J", "D", "M", "T", "S", "B"})
	state[5].PushMultiple([]string{"N", "C", "D", "R", "J"})
	state[6].PushMultiple([]string{"W", "L", "D", "Q", "P", "J", "G", "Z"})
	state[7].PushMultiple([]string{"P", "Z", "T", "F", "R", "H"})
	state[8].PushMultiple([]string{"L", "V", "M", "G"})
	state[9].PushMultiple([]string{"C", "B", "G", "P", "F", "Q", "R", "J"})

	return state
}

func getInstructions(sessionCookie string) []*Instruction {
	allRows := util.GetRows("https://adventofcode.com/2022/day/5/input", sessionCookie)
	instuctionRows := allRows[10:]

	result := make([]*Instruction, len(instuctionRows)-1)
	for i, instruction := range instuctionRows {
		if instruction == "" {
			continue
		}
		result[i] = newInstruction(instruction)
	}

	return result
}

func newInstruction(instructionText string) *Instruction {
	parts := strings.Split(instructionText, " ")
	if len(parts) != 6 {
		fmt.Printf("InstructionText: '%s' doesn't have enough parts to compose instruction\n", instructionText)
		return nil
	}
	move, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("InstructionText: '%s' parts[1] err: %v\n", instructionText, err)
		return nil
	}
	from, err := strconv.Atoi(parts[3])
	if err != nil {
		fmt.Printf("InstructionText: '%s' parts[3] err: %v\n", instructionText, err)
		return nil
	}
	to, err := strconv.Atoi(parts[5])
	if err != nil {
		fmt.Printf("InstructionText: '%s' parts[5] err: %v\n", instructionText, err)
		return nil
	}

	return &Instruction{
		Move: move,
		From: from,
		To:   to,
	}
}
