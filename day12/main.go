package main

import (
	"aoc2022/utils"
	"math"
	"strings"
)

func main() {
	utils.Run("day12", functionPart1, functionPart2)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

type AdjencyMatrix [][]float64
type Graph struct {
	adjencyMatrix AdjencyMatrix
	starts        []int
	end           int
}

func canClimb(lhs, rhs string) bool {
	return strings.Index(alphabet, lhs) <= strings.Index(alphabet, rhs)+1
}

func getGraph(input string, includeA bool) Graph {
	graph := Graph{}
	graph.adjencyMatrix = [][]float64{}

	lines := strings.Split(input, "\n")
	line := strings.Join(lines, "")
	vertices := len(line)
	for i := 0; i < vertices; i++ {
		graph.adjencyMatrix = append(graph.adjencyMatrix, make([]float64, vertices))
		for j := 0; j < vertices; j++ {
			value := math.Inf(0)
			graph.adjencyMatrix[i][j] = value
		}
	}

	for i := 0; i < vertices; i++ {
		currentValue := line[i]
		if currentValue == 'S' {
			graph.starts = append(graph.starts, i)
		} else if currentValue == 'E' {
			graph.end = i
		} else if currentValue == 'a' && includeA {
			graph.starts = append(graph.starts, i)
		}

		currentGraph := graph.adjencyMatrix[i]
		// Get top of current index in line
		topIdx := i - len(lines[0])
		rightIdx := i + 1
		bottomIdx := i + len(lines[0])
		leftIdx := i - 1
		if topIdx >= 0 {
			topValue := line[topIdx]
			if canClimb(string(topValue), string(currentValue)) {
				currentGraph[topIdx] = 1
			}

		}
		if rightIdx < vertices {
			rightValue := line[rightIdx]
			if canClimb(string(rightValue), string(currentValue)) {
				currentGraph[rightIdx] = 1
			}
		}
		if bottomIdx < vertices {
			bottomValue := line[bottomIdx]
			if canClimb(string(bottomValue), string(currentValue)) {
				currentGraph[bottomIdx] = 1
			}
		}
		if leftIdx >= 0 {
			leftValue := line[leftIdx]
			if canClimb(string(leftValue), string(currentValue)) {
				currentGraph[leftIdx] = 1
			}
		}
	}
	return graph
}

func minDistance(distances []float64, sptSet []bool) int {
	min := math.Inf(0)
	minIndex := -1

	for v := 0; v < len(distances); v++ {
		if !sptSet[v] && distances[v] <= min {
			min = distances[v]
			minIndex = v
		}
	}

	return minIndex
}

func dijkstra(graph Graph, start int) int {
	distances := make([]float64, len(graph.adjencyMatrix))
	sptSet := make([]bool, len(graph.adjencyMatrix))

	for i := 0; i < len(graph.adjencyMatrix); i++ {
		distances[i] = math.Inf(0)
		sptSet[i] = false
	}

	distances[start] = 1
	for i := 0; i < len(graph.adjencyMatrix)-1; i++ {
		u := minDistance(distances, sptSet)
		sptSet[u] = true
		for v := 0; v < len(graph.adjencyMatrix); v++ {
			if !sptSet[v] && graph.adjencyMatrix[u][v] != 0 && distances[u] != math.MaxInt32 && distances[u]+graph.adjencyMatrix[u][v] < distances[v] {
				distances[v] = distances[u] + graph.adjencyMatrix[u][v]
			}
		}
	}

	return int(distances[graph.end-1])
}

func functionPart1(input string) int {
	graph := getGraph(input, false)
	return dijkstra(graph, graph.starts[0])
}

func functionPart2(input string) int {
	graph := getGraph(input, true)
	// Get min distance
	min := math.MaxInt
	for _, start := range graph.starts {
		distance := dijkstra(graph, start)
		if distance < min && distance != math.MinInt {
			min = distance
		}
	}
	return min
}
