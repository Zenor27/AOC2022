package part1

import "strings"

type Shape int64
type Score int64

const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

const (
	Lose Score = 0
	Draw Score = 3
	Win  Score = 6
)

type Round struct {
	opponentPlay Shape
	myPlay       Shape
}


func getShape(strShape string) Shape {
	switch strShape {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	panic("Unknown shape")
}

func getRound(strRound string) Round {
	round := strings.Split(strRound, " ")
	opponentPlayStr := round[0]
	myPlayStr := round[1]
	return Round{
		opponentPlay: getShape(opponentPlayStr),
		myPlay:       getShape(myPlayStr),
	}
}

func getRounds(input string) []Round {
	strRounds := strings.Split(input, "\n")
	rounds := make([]Round, 0, len(strRounds))
	for _, strRound := range strRounds {
		rounds = append(rounds, getRound(strRound))
	}
	return rounds
}

func getScore(round Round) int {
	if round.myPlay == round.opponentPlay {
		return int(Draw)
	}
	switch round.opponentPlay {
	case Rock:
		if (round.myPlay == Paper) {
			return int(Win)
		}
	case Paper:
		if (round.myPlay == Scissors) {
			return int(Win)
		}
	case Scissors:
		if (round.myPlay == Rock) {
			return int(Win)
		}
	}
	return int(Lose)
}

func FunctionPart1(input string) int {
	rounds := getRounds(input)

	score := 0
	for _, round := range rounds {
		score += int(round.myPlay)
		score += getScore(round)
	}

	return score
}