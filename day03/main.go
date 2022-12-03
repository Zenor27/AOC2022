package main

import (
	"aoc2022/utils"
	"strings"
)

func main() {
	utils.Run("day03", functionPart1, functionPart2)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func getCommonLetters(rucksacks []string) []string {
	letters := make([]string, 0)
	for _, rucksack := range rucksacks {
		middleIndex := len(rucksack) / 2
		// Split rucksack into two parts
		leftCompartment := rucksack[:middleIndex]
		rightCompartment := rucksack[middleIndex:]

		// Common letter in both compartments
		commonLetter := ""
		for _, letter := range leftCompartment {
			if strings.Contains(rightCompartment, string(letter)) {
				commonLetter = string(letter)
				break
			}
		}
		letters = append(letters, commonLetter)
	}
	return letters
}

func getScore(commonLetters []string) int {
	score := 0
	for _, commonLetter := range commonLetters {
		index := strings.Index(alphabet, commonLetter) + 1

		// Uppercase
		if index == 0 {
			letter := strings.ToLower(commonLetter)
			index = strings.Index(alphabet, letter) + 1
			score += index + 26
		} else {
			score += index
		}
	}
	return score
}

func functionPart1(input string) int {
	rucksacks := strings.Split(input, "\n")
	commonLetters := getCommonLetters(rucksacks)
	score := getScore(commonLetters)
	return score
}

func getGroupedRucksacks(rucksacks []string) [][]string {
	// Group rucksacks by 3
	groupedRucksacks := make([][]string, 0)
	for i := 0; i < len(rucksacks); i += 3 {
		groupedRucksacks = append(groupedRucksacks, rucksacks[i:i+3])
	}
	return groupedRucksacks
}

func functionPart2(input string) int {
	rucksacks := strings.Split(input, "\n")
	groupedRucksacks := getGroupedRucksacks(rucksacks)

	commonLetters := make([]string, 0)
	for _, groupedRucksack := range groupedRucksacks {
		firstRucksack, secondRucksack, thirdRucksack := groupedRucksack[0], groupedRucksack[1], groupedRucksack[2]

		commonLetter := ""
		for _, letter := range firstRucksack {
			if strings.Contains(secondRucksack, string(letter)) && strings.Contains(thirdRucksack, string(letter)) {
				commonLetter = string(letter)
				break
			}
		}

		commonLetters = append(commonLetters, commonLetter)
	}

	return getScore(commonLetters)
}
