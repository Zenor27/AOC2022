package main

import (
	"aoc2022/utils"
)

func main() {
	utils.Run("day06", functionPart1, functionPart2)
}

func functionPart1(input string) int {
	for i := 0; i < len(input)-3; i++ {
		start, end := i, i+4
		group := input[start:end]
		if utils.HasOnlyUniqChars(group) {
			return end
		}
	}
	return -1
}

func functionPart2(input string) int {
	for i := 0; i < len(input)-13; i++ {
		start, end := i, i+14
		group := input[start:end]
		if utils.HasOnlyUniqChars(group) {
			return end
		}
	}
	return -1
}
