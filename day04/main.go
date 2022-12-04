package main

import (
	"aoc2022/utils"
	"strconv"
	"strings"
)

func main() {
	utils.Run("day04", functionPart1, functionPart2)
}

func getRangeFromElf(elf string) (int, int) {
	elfRange := strings.Split(elf, "-")
	elfMin, _ := strconv.Atoi(elfRange[0])
	elfMax, _ := strconv.Atoi(elfRange[1])
	return elfMin, elfMax
}

func functionPart1(input string) int {
	pairs := strings.Split(input, "\n")
	ovelaps := 0
	for _, pair := range pairs {
		elves := strings.Split(pair, ",")
		elf1, elf2 := elves[0], elves[1]
		elf1Min, elf1Max := getRangeFromElf(elf1)
		elf2Min, elf2Max := getRangeFromElf(elf2)

		if (elf1Min <= elf2Min && elf1Max >= elf2Max) || (elf2Min <= elf1Min && elf2Max >= elf1Max) {
			ovelaps++
		}
	}
	return ovelaps
}

func functionPart2(input string) int {
	pairs := strings.Split(input, "\n")
	ovelaps := 0
	for _, pair := range pairs {
		elves := strings.Split(pair, ",")
		elf1, elf2 := elves[0], elves[1]
		elf1Min, elf1Max := getRangeFromElf(elf1)
		elf2Min, elf2Max := getRangeFromElf(elf2)

		if elf1Max >= elf2Min && elf1Min <= elf2Max {
			ovelaps++
		}
	}
	return ovelaps
}
