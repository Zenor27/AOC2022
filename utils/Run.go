package utils

import (
	"flag"
	"fmt"
)

func printOutput(output int) {
	fmt.Println("Output:")
	fmt.Println(output)
}

func Run(currentDirectory string, functionPart1 func(input string) int, functionPart2 ...func(input string) int) {
	part := flag.Int("part", 1, "Part 1 or 2")
	flag.Parse()
	if part == nil {
		panic("part is nil")
	} else if *part < 1 || *part > 2 {
		panic("part argument must be 1 or 2")
	} else if *part == 2 && len(functionPart2) == 0 {
		panic("part 2 function is missing")
	}

	inputPath := fmt.Sprintf("%s/input_part%d.txt", currentDirectory, *part)
	input := ReadFile(inputPath)

	if *part == 1 {
		printOutput(functionPart1(input))
		return
	}

	printOutput(functionPart2[0](input))
}
