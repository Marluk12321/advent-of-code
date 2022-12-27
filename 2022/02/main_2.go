package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Choice int

const (
	ROCK Choice = iota
	PAPER
	SCISSORS
)

var strToChoice = map[string]Choice{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var choiceToScore = map[Choice]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

type Outcome int

const (
	LOSS Outcome = iota
	DRAW
	WIN
)

var strToOutcome = map[string]Outcome{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}

var outcomeToScore = map[Outcome]int{
	LOSS: 0,
	DRAW: 3,
	WIN:  6,
}

func getMyChoice(other_choice Choice, needed_outcome Outcome) Choice {
	if needed_outcome == DRAW {
		return other_choice
	}
	if needed_outcome == WIN {
		switch other_choice {
		case ROCK:
			return PAPER
		case PAPER:
			return SCISSORS
		default:
			return ROCK
		}
	} else {
		switch other_choice {
		case ROCK:
			return SCISSORS
		case PAPER:
			return ROCK
		default:
			return PAPER
		}
	}
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Missing file path argument")
		return
	}

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open", path)
		fmt.Println(err)
		return
	}
	defer f.Close()

	var totalScore int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Fields(line)
		other_choice := strToChoice[choices[0]]
		needed_outcome := strToOutcome[choices[1]]

		totalScore += outcomeToScore[needed_outcome]
		my_choice := getMyChoice(other_choice, needed_outcome)
		totalScore += choiceToScore[my_choice]
		fmt.Println(line, my_choice)
	}

	fmt.Println("Total score:", totalScore)
}
