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
	"A": ROCK, "X": ROCK,
	"B": PAPER, "Y": PAPER,
	"C": SCISSORS, "Z": SCISSORS,
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

var outcomeToScore = map[Outcome]int{
	LOSS: 0,
	DRAW: 3,
	WIN:  6,
}

func getOutcome(other_choice Choice, my_choice Choice) Outcome {
	if other_choice == my_choice {
		return DRAW
	}
	switch other_choice {
	case ROCK:
		if my_choice == PAPER {
			return WIN
		} else {
			return LOSS
		}
	case PAPER:
		if my_choice == SCISSORS {
			return WIN
		} else {
			return LOSS
		}
	default:
		if my_choice == ROCK {
			return WIN
		} else {
			return LOSS
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
		my_choice := strToChoice[choices[1]]

		totalScore += choiceToScore[my_choice]
		outcome := getOutcome(other_choice, my_choice)
		totalScore += outcomeToScore[outcome]
		fmt.Println(line, outcome)
	}

	fmt.Println("Total score:", totalScore)
}
