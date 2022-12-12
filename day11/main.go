package main

import (
	"aoc2022/utils"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/samber/lo"
)

type Monkey struct {
	id                 int
	itemWorries        []int
	op                 func(int) int
	divisibleBy        int
	test               func(int) bool
	trueThrowToMonkey  int
	falseThrowToMonkey int
}

func main() {
	utils.Run("day11", functionPart1, functionPart2)
}

func getIdToMonkey(input string) *treemap.Map {
	idToMonkey := treemap.NewWithIntComparator()

	for _, monkeyLines := range utils.SplitByEmptyLine(input) {
		monkeyLine := strings.Split(monkeyLines, "\n")

		monkeyLineId := monkeyLine[0]
		monkeyId := utils.ToInt(strings.Split(strings.Split(monkeyLineId, " ")[1], ":")[0])

		startingItemsLine := monkeyLine[1]
		itemWorriesStr := strings.Split(strings.Split(startingItemsLine, ": ")[1], ", ")
		itemWorries := lo.Map(itemWorriesStr, func(itemWorryStr string, _ int) int {
			return utils.ToInt(itemWorryStr)
		})

		operationLine := monkeyLine[2]
		operation := strings.Split(strings.Split(operationLine, "= ")[1], " ")
		operand, rhs := operation[1], operation[2]

		testLine := monkeyLine[3]
		divisibleBy := utils.ToInt(strings.Split(testLine, "by ")[1])

		trueLine := monkeyLine[4]
		trueThrowToMonkey := utils.ToInt(strings.Split(trueLine, "to monkey ")[1])
		falseLine := monkeyLine[5]
		falseThrowToMonkey := utils.ToInt(strings.Split(falseLine, "to monkey ")[1])

		monkey := Monkey{
			id:          monkeyId,
			itemWorries: itemWorries,
			op: func(lhs int) int {
				rhs_ := lhs
				if rhs != "old" {
					rhs_ = int(utils.ToInt(rhs))
				}

				switch operand {
				case "+":
					return lhs + rhs_
				case "-":
					return lhs - rhs_
				case "*":
					return lhs * rhs_
				case "/":
					return lhs / rhs_
				}
				panic("Unknown operand: " + operand)
			},
			divisibleBy: divisibleBy,
			test: func(lhs int) bool {
				return lhs%divisibleBy == 0
			},
			trueThrowToMonkey:  trueThrowToMonkey,
			falseThrowToMonkey: falseThrowToMonkey,
		}
		idToMonkey.Put(monkeyId, &monkey)
	}

	return idToMonkey
}

func getResult(idToInspected map[int]int) int {
	// Get two max inspected monkey
	max1, max2 := 0, 0
	for _, inspectCount := range idToInspected {
		if inspectCount > max1 {
			max2 = max1
			max1 = inspectCount
		} else if inspectCount > max2 {
			max2 = inspectCount
		}
	}

	return max1 * max2
}

func functionPart1(input string) int {
	idToMonkey := getIdToMonkey(input)

	idToInspected := make(map[int]int)
	for round := 0; round < 20; round++ {
		for _, monkey := range idToMonkey.Values() {
			monkey := monkey.(*Monkey)
			for _, itemWorry := range monkey.itemWorries {
				idToInspected[monkey.id]++
				worryLevel := monkey.op(itemWorry)
				worryLevel /= 3
				monkeyToThrow := lo.TernaryF(monkey.test(worryLevel), func() *Monkey {
					x, _ := idToMonkey.Get(monkey.trueThrowToMonkey)
					return x.(*Monkey)
				}, func() *Monkey {
					x, _ := idToMonkey.Get(monkey.falseThrowToMonkey)
					return x.(*Monkey)
				})

				monkeyToThrow.itemWorries = append(monkeyToThrow.itemWorries, worryLevel)
				// Item is thrown away, remove from current monkey
				monkey.itemWorries = monkey.itemWorries[:len(monkey.itemWorries)-1]
			}
		}
	}
	return getResult(idToInspected)
}

func functionPart2(input string) int {
	idToMonkey := getIdToMonkey(input)

	idToInspected := make(map[int]int)
	commonDivisor := 1
	for _, monkey := range idToMonkey.Values() {
		monkey := monkey.(*Monkey)
		commonDivisor *= monkey.divisibleBy
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range idToMonkey.Values() {
			monkey := monkey.(*Monkey)
			idToInspected[monkey.id] += len(monkey.itemWorries)
			for _, itemWorry := range monkey.itemWorries {
				worryLevel := monkey.op(itemWorry)
				monkeyToThrow := lo.TernaryF(monkey.test(worryLevel), func() *Monkey {
					x, _ := idToMonkey.Get(monkey.trueThrowToMonkey)
					return x.(*Monkey)
				}, func() *Monkey {
					x, _ := idToMonkey.Get(monkey.falseThrowToMonkey)
					return x.(*Monkey)
				})

				monkeyToThrow.itemWorries = append(monkeyToThrow.itemWorries, worryLevel%commonDivisor)
			}
			// Item is thrown away, clean up
			monkey.itemWorries = []int{}
		}
	}
	return getResult(idToInspected)
}
