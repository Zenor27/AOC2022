package part2

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
	shape Shape
	score Score
}

func getShape(opponentPlayStr string, score Score) Shape {
	switch opponentPlayStr {
	case "A":
		if score == Win {
			return Paper
		} else if score == Lose {
			return Scissors
		}
		return Rock
	case "B":
		if score == Win {
			return Scissors
		} else if score == Lose {
			return Rock
		}
		return Paper
	case "C":
		if score == Win {
			return Rock
		} else if score == Lose {
			return Paper
		}
		return Scissors
	}
	panic("Unknown shape")
}

func getScore(strScore string) Score {
	switch strScore {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	panic("Unknown score")
}

func getRound(strRound string) Round {
	round := strings.Split(strRound, " ")
	opponentPlayStr := round[0]
	myPlayStr := round[1]
	score := getScore(myPlayStr)
	return Round{
		shape: getShape(opponentPlayStr, score),
		score: score,
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

func FunctionPart2(input string) int {
	rounds := getRounds(input)

	score := 0
	for _, round := range rounds {
		score += int(round.score)
		score += int(round.shape)
	}

	return score
}
