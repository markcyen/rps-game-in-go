package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	// PLAYERWINS   = 1
	// COMPUTERWINS = 2
	// DRAW         = 3
)

var playerWinningNote = []string{
	"Time to buy a lottery!",
	"You ROCK (or PAPER or SCISSORS)!",
	"You are such a winner!",
}

var computerWinningNote = []string{
	"Try again!",
	"Don't let the computer beat you!",
	"Awwwh, nice try and try again!",
}

var drawNote = []string{
	"Great minds!",
	"Deuce!",
	"Try again and again!",
}

type Round struct {
	ResultMessage  string `json:"result_message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	resultMessage := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		resultMessage = drawNote[rand.Intn(len(drawNote))]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		resultMessage = playerWinningNote[rand.Intn((len(playerWinningNote)))]
	} else {
		roundResult = "Computer wins!"
		resultMessage = computerWinningNote[rand.Intn((len(computerWinningNote)))]
	}

	// return winner, computerChoice, roundResult

	var result Round
	result.ResultMessage = resultMessage
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
