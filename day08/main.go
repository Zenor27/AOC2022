package main

import (
	"aoc2022/utils"
	"strconv"
	"strings"
)

func main() {
	utils.Run("day08", functionPart1, functionPart2)
}

type Tree struct {
	height int
}

func sumTop(forest [][]Tree, visibleTrees, nonBlockingTrees *int, currentHeight, i, j int) bool {
	visible := true
	for k := i - 1; k >= 0; k-- {
		if forest[k][j].height >= currentHeight {
			visible = false
			break
		}
		*nonBlockingTrees++
	}
	if visible {
		*visibleTrees++
	}
	return visible
}

func sumBottom(forest [][]Tree, visibleTrees, nonBlockingTrees *int, currentHeight, i, j int) bool {
	visible := true
	for k := i + 1; k < len(forest); k++ {
		if forest[k][j].height >= currentHeight {
			visible = false
			break
		}
		*nonBlockingTrees++
	}
	if visible {
		*visibleTrees++
	}
	return visible
}

func sumLeft(forest [][]Tree, visibleTrees, nonBlockingTrees *int, currentHeight, i, j int) bool {
	visible := true
	for k := j - 1; k >= 0; k-- {
		if forest[i][k].height >= currentHeight {
			visible = false
			break
		}
		*nonBlockingTrees++
	}
	if visible {
		*visibleTrees++
	}
	return visible
}

func sumRight(forest [][]Tree, visibleTrees, nonBlockingTrees *int, currentHeight, i, j int) bool {
	visible := true
	for k := j + 1; k < len(forest[i]); k++ {
		if forest[i][k].height >= currentHeight {
			visible = false
			break
		}
		*nonBlockingTrees++
	}
	if visible {
		*visibleTrees++
	}
	return visible
}

func getForest(input string) [][]Tree {
	var forest [][]Tree

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trees := make([]Tree, 0)
		for _, char := range line {
			height, _ := strconv.Atoi(string(char))
			trees = append(trees, Tree{height: height})
		}
		forest = append(forest, trees)
	}
	return forest
}

func functionPart1(input string) int {
	forest := getForest(input)

	edgeTrees := 0
	visibleTrees := 0
	nonBlockingTrees := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			// If tree is on edge, it is always visible
			if i == 0 || j == 0 || i == len(forest)-1 || j == len(forest[i])-1 {
				edgeTrees++
				continue
			}

			currentHeight := forest[i][j].height
			// A tree is visible if all of the other trees between it and an edge of the grid are shorter than it
			visible := sumTop(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if visible {
				continue
			}
			visible = sumBottom(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if visible {
				continue
			}
			visible = sumLeft(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if visible {
				continue
			}
			sumRight(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)

		}
	}

	return visibleTrees + edgeTrees
}

func functionPart2(input string) int {
	forest := getForest(input)

	scenicScores := make([]int, 0)
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			// If tree is on edge, it is always visible
			if i == 0 || j == 0 || i == len(forest)-1 || j == len(forest[i])-1 {
				scenicScores = append(scenicScores, 0)
				continue
			}

			currentHeight := forest[i][j].height
			visibleTrees := 0
			nonBlockingTrees := 0

			scenicScore := 1
			valid := sumTop(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if !valid {
				nonBlockingTrees++
			}
			scenicScore *= nonBlockingTrees

			nonBlockingTrees = 0
			valid = sumBottom(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if !valid {
				nonBlockingTrees++
			}
			scenicScore *= nonBlockingTrees

			nonBlockingTrees = 0
			valid = sumLeft(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if !valid {
				nonBlockingTrees++
			}
			scenicScore *= nonBlockingTrees

			nonBlockingTrees = 0
			valid = sumRight(forest, &visibleTrees, &nonBlockingTrees, currentHeight, i, j)
			if !valid {
				nonBlockingTrees++
			}
			scenicScore *= nonBlockingTrees

			scenicScores = append(scenicScores, scenicScore)
		}
	}

	return utils.GetMaxFromList(scenicScores)
}
