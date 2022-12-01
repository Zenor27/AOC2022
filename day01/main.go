package main

import (
	"aoc2022/utils"
	"sort"
	"strconv"
	"strings"
)

func main() {
	utils.Run("day01", functionPart1, functionPart2)
}

func sumElfCalories(elfCalories string) int {
	calories := strings.Split(elfCalories, "\n")
	elfCaloriesInt := 0
	for _, calorie := range calories {
		calorieInt, _ := strconv.ParseInt(calorie, 10, 64)
		elfCaloriesInt += int(calorieInt)
	}

	return elfCaloriesInt
}

func summedCaloriesPerElfes(caloriesPerElfes []string) []int {
	summedCaloriesPerElfes := make([]int, len(caloriesPerElfes))
	for _, elfCalories := range caloriesPerElfes {
		summedCaloriesPerElfes = append(summedCaloriesPerElfes, sumElfCalories(elfCalories))
	}
	return summedCaloriesPerElfes

}

func functionPart1(input string) int {
	caloriesPerElfes := utils.SplitByEmptyLine(input)
	summedCaloriesPerElfes := summedCaloriesPerElfes(caloriesPerElfes)
	sort.Slice(summedCaloriesPerElfes, func(i, j int) bool {
		return summedCaloriesPerElfes[i] > summedCaloriesPerElfes[j]
	})
	return summedCaloriesPerElfes[0]
}

func functionPart2(input string) int {
	caloriesPerElfes := utils.SplitByEmptyLine(input)
	summedCaloriesPerElfes := summedCaloriesPerElfes(caloriesPerElfes)

	sort.Slice(summedCaloriesPerElfes, func(i, j int) bool {
		return summedCaloriesPerElfes[i] > summedCaloriesPerElfes[j]
	})

	return summedCaloriesPerElfes[0] + summedCaloriesPerElfes[1] + summedCaloriesPerElfes[2]
}
