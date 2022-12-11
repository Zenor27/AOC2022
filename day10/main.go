package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	utils.Run("day10", functionPart1, functionPart2)
}

type Op string

const (
	NOOP = "noop"
	ADDX = "addx"
)

type Instruction struct {
	op    Op
	value int
}

func getInstructions(input string) []Instruction {
	instructions := make([]Instruction, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			instructions = append(instructions, Instruction{NOOP, 0})
		} else {
			x := strings.Split(line, " ")
			value, _ := strconv.Atoi(x[1])
			instructions = append(instructions, Instruction{ADDX, value})
		}
	}
	return instructions
}

func functionPart1(input string) int {
	instructions := getInstructions(input)
	cycleToInstructions := make(map[int]Instruction)

	X := 1
	currentCycle := 1
	strength := 0
	nextCycle := currentCycle + 2
	for {
		if len(instructions) != 0 {
			instruction := instructions[0]
			instructions = instructions[1:]
			if instruction.op == ADDX {
				cycleToInstructions[nextCycle] = instruction
				nextCycle = nextCycle + 2
			} else {
				nextCycle++
			}
		}

		currentCycle++

		if instructionToExecute, ok := cycleToInstructions[currentCycle]; ok {
			X += instructionToExecute.value
			delete(cycleToInstructions, currentCycle)
		}

		switch currentCycle {
		case 20, 60, 100, 140, 180, 220:
			strength += (X * currentCycle)
		}

		if len(instructions) == 0 && len(cycleToInstructions) == 0 {
			break
		}
	}

	return strength
}

func functionPart2(input string) int {
	instructions := getInstructions(input)
	cycleToInstructions := make(map[int]Instruction)

	X := 1
	currentCycle := 1
	nextCycle := currentCycle + 2
	currentPixelPosition := 0
	for {
		if len(instructions) != 0 {
			instruction := instructions[0]
			instructions = instructions[1:]
			if instruction.op == ADDX {
				cycleToInstructions[nextCycle] = instruction
				nextCycle = nextCycle + 2
			} else {
				nextCycle++
			}
		}

		currentPixelPosition++

		if currentPixelPosition == X || currentPixelPosition == X+1 || currentPixelPosition == X+2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		currentCycle++
		if instructionToExecute, ok := cycleToInstructions[currentCycle]; ok {
			X += instructionToExecute.value
			delete(cycleToInstructions, currentCycle)
		}

		switch currentCycle - 1 {
		case 40, 80, 120, 160, 200, 240:
			currentPixelPosition = 0
			fmt.Println()
		}

		if len(instructions) == 0 && len(cycleToInstructions) == 0 && currentPixelPosition == 0 {
			fmt.Println()
			break
		}
	}

	return 0
}
