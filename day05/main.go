package main

import (
	"aoc2022/utils"
	"math"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/stacks"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func main() {
	utils.Run("day05", functionPart1, functionPart2)
}

type Move struct {
	from int
	to   int
	pop  int
}

func parseMoves(input string) []Move {
	inputSplit := utils.SplitByEmptyLine(input)
	inputMoves := inputSplit[len(inputSplit)-1]
	moves := make([]Move, 0)
	for _, inputMove := range strings.Split(inputMoves, "\n") {
		splitInputMove := strings.Split(inputMove, " ")
		pop, _ := strconv.Atoi(splitInputMove[1])
		from, _ := strconv.Atoi(splitInputMove[3])
		to, _ := strconv.Atoi(splitInputMove[5])
		move := Move{
			pop:  pop,
			from: from - 1,
			to:   to - 1,
		}
		moves = append(moves, move)
	}
	return moves
}

func parseCratesStacks(input string) []stacks.Stack {
	inputSplit := utils.SplitByEmptyLine(input)
	inputCratesStacks := strings.Split(inputSplit[0], "\n")
	nbCratesStacks := math.Ceil(float64(len(inputCratesStacks[0])) / 4)

	cratesStacks := make([]stacks.Stack, 0)
	for i := 0; i < int(nbCratesStacks); i++ {
		cratesStacks = append(cratesStacks, arraystack.New())
	}

	// Remove last line
	inputCratesStacks = inputCratesStacks[:len(inputCratesStacks)-1]

	// Reverse inputCratesStacks
	for i := len(inputCratesStacks)/2 - 1; i >= 0; i-- {
		opp := len(inputCratesStacks) - 1 - i
		inputCratesStacks[i], inputCratesStacks[opp] = inputCratesStacks[opp], inputCratesStacks[i]
	}

	for _, inputCratesStack := range inputCratesStacks {
		for idx, cratesStack := range cratesStacks {
			// Offset is for: idx = 0 -> 1, idx = 1 -> 5, idx = 2 -> 9, ...
			offset := (idx * 4) + 1
			crateId := inputCratesStack[offset]
			if crateId != ' ' {
				cratesStack.Push(string(crateId))
			}
		}
	}

	return cratesStacks
}

func parseInput(input string) ([]stacks.Stack, []Move) {
	moves := parseMoves(input)
	cratesStacks := parseCratesStacks(input)

	return cratesStacks, moves
}

func applyMove(cratesStacks []stacks.Stack, move Move) {
	for i := 0; i < move.pop; i++ {
		crates := cratesStacks[move.from]
		crate, _ := crates.Pop()
		cratesStacks[move.to].Push(crate)
	}
}

func getTops(cratesStacks []stacks.Stack) []string {
	tops := make([]string, 0)
	for _, crates := range cratesStacks {
		topCreate, _ := crates.Peek()
		tops = append(tops, topCreate.(string))
	}
	return tops
}

func functionPart1(input string) string {
	cratesStacks, moves := parseInput(input)
	for _, move := range moves {
		applyMove(cratesStacks, move)
	}
	return strings.Join(getTops(cratesStacks), "")
}

func applyMovePart2(cratesStacks []stacks.Stack, move Move) {
	tmpStack := arraystack.New()

	for i := 0; i < move.pop; i++ {
		crates := cratesStacks[move.from]
		crate, _ := crates.Pop()
		tmpStack.Push(crate)
	}
	it := tmpStack.Iterator()
	for it.Next() {
		cratesStacks[move.to].Push(it.Value())
	}
}

func functionPart2(input string) string {
	cratesStacks, moves := parseInput(input)
	for _, move := range moves {
		applyMovePart2(cratesStacks, move)
	}
	return strings.Join(getTops(cratesStacks), "")
}
