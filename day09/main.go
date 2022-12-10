package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	utils.Run("day09", functionPart1, functionPart2)
}

type Motion struct {
	direction string
	steps     int
}

func getMotions(input string) []Motion {
	lines := strings.Split(input, "\n")

	motions := make([]Motion, len(lines))
	for i, line := range lines {
		steps, _ := strconv.Atoi(line[2:])
		motions[i] = Motion{line[0:1], steps}
	}
	return motions
}

func functionPart1(input string) int {
	motions := getMotions(input)

	headX := 0
	headY := 0
	tailX := 0
	tailY := 0
	// Create a set of tuples to store the visited points
	visited := make(map[string]bool)
	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {
			switch motion.direction {
			case "R":
				headX++
			case "L":
				headX--
			case "U":
				headY++
			case "D":
				headY--
			}

			// Check if the tail is next to head (diagonal included)
			isTouched := false
			for _, x := range []int{-1, 0, 1} {
				for _, y := range []int{-1, 0, 1} {
					if headX+x == tailX && headY+y == tailY {
						isTouched = true
						break
					}
				}
				if isTouched {
					break
				}
			}

			if !isTouched {
				switch motion.direction {
				case "R":
					tailX++
					if headY != tailY {
						tailY = headY
					}
				case "L":
					tailX--
					if headY != tailY {
						tailY = headY
					}
				case "U":
					tailY++
					if headX != tailX {
						tailX = headX
					}
				case "D":
					tailY--
					if headX != tailX {
						tailX = headX
					}
				}
			}

			// Add tail to visited
			visited[fmt.Sprintf("%d,%d", tailX, tailY)] = true
		}
	}

	return len(visited)
}

type Knot struct {
	x        int
	y        int
	children *Knot
}

func getKnots() []*Knot {
	knots := make([]*Knot, 10)
	for i := 0; i < 10; i++ {
		knots[i] = &Knot{0, 0, nil}
	}

	for i := 0; i < 9; i++ {
		knots[i].children = knots[i+1]
	}

	return knots
}

func move(knot *Knot, motion Motion) {
	switch motion.direction {
	case "R":
		(*knot).x++
	case "L":
		(*knot).x--
	case "U":
		(*knot).y++
	case "D":
		(*knot).y--
	}
}

func offset(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

func moveTail(parent, child *Knot, motion Motion) {
	dX := (*parent).x - (*child).x
	dY := (*parent).y - (*child).y

	if math.Abs(float64(dX)) <= 1 && math.Abs(float64(dY)) <= 1 {
		return
	}

	(*child).x += offset(dX)
	(*child).y += offset(dY)
}

func functionPart2(input string) int {
	motions := getMotions(input)
	knots := getKnots()

	// Create a set of tuples to store the visited points
	visited := make(map[string]bool)

	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {
			move(knots[0], motion)
			for i := 1; i < len(knots); i++ {
				moveTail(knots[i-1], knots[i], motion)
			}

			tail := knots[len(knots)-1]
			visited[fmt.Sprintf("%d,%d", tail.x, tail.y)] = true

			continue
		}
	}

	return len(visited)
}
