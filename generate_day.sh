#!/bin/sh

# Check if there is exactly one argument
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

DAY=$1
# If day is < 10
if [ $DAY -lt 10 ]; then
    DAY="0$DAY"
fi

echo "🚀 Generating day $DAY"

# Check if day already exists
if [ -d "day$DAY" ]; then
    echo "⚠️ Day $DAY already exists"
    exit 1
fi

mkdir -p "day$DAY"
touch "day$DAY/input_part1.txt"
touch "day$DAY/input_part2.txt"
touch "day$DAY/main.go"


# File content here
CONTENT="
package main

import (
	\"aoc2022/utils\"
)

func main() {
	utils.Run(\"day$DAY\", functionPart1, functionPart2)
}

func functionPart1(input string) int {
	return 0
}

func functionPart2(input string) int {
	return 0
}
"

echo "$CONTENT" > "day$DAY/main.go"

