package main

import (
	"aoc2022/utils"
	"strings"

	"encoding/json"
)

func main() {
	utils.Run("day13", functionPart1, functionPart2)
}

func compare(a, b any) float64 {
	aArray, aOk := a.([]any)
	bArray, bOk := b.([]any)

	if !aOk && !bOk {
		return a.(float64) - b.(float64)
	}

	if !aOk {
		aArray = []any{a}
	}
	if !bOk {
		bArray = []any{b}
	}

	for i := 0; i < len(aArray) && i < len(bArray); i++ {
		if c := compare(aArray[i], bArray[i]); c != 0 {
			return c
		}
	}

	return float64(len(aArray) - len(bArray))
}

func functionPart1(input string) int {
	pairs := utils.SplitByEmptyLine(input)
	sortedPairs := 0
	for idx, pair := range pairs {
		split := strings.Split(pair, "\n")
		var packet1, packet2 any
		json.Unmarshal([]byte(split[0]), &packet1)
		json.Unmarshal([]byte(split[1]), &packet2)

		if c := compare(packet1, packet2); c < 0 {
			sortedPairs += (idx + 1)
		}
	}
	return sortedPairs
}

const dividerPacket1 = "[[2]]"
const dividerPacket2 = "[[6]]"
const dividerPackets = "[[2]]\n[[6]]"

func functionPart2(input string) int {
	pairs := utils.SplitByEmptyLine(input)
	pairs = append(pairs, dividerPackets)

	allPackets := make([]any, 0)
	for _, pairs := range pairs {
		split := strings.Split(pairs, "\n")
		var packet1, packet2 any
		json.Unmarshal([]byte(split[0]), &packet1)
		json.Unmarshal([]byte(split[1]), &packet2)
		allPackets = append(allPackets, packet1, packet2)
	}

	// Sort all packets by compare function
	for i := 0; i < len(allPackets); i++ {
		for j := i + 1; j < len(allPackets); j++ {
			if c := compare(allPackets[i], allPackets[j]); c > 0 {
				allPackets[i], allPackets[j] = allPackets[j], allPackets[i]
			}
		}
	}

	// Find dividerPacket1 idx and dividerPacket2 idx
	dividerPacket1Idx := -1
	dividerPacket2Idx := -1
	for idx, packet := range allPackets {
		x, _ := json.Marshal(packet)
		if string(x) == dividerPacket1 {
			dividerPacket1Idx = idx + 1
		} else if string(x) == dividerPacket2 {
			dividerPacket2Idx = idx + 1
		}
	}

	return dividerPacket1Idx * dividerPacket2Idx
}
