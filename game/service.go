package game

import (
	"container/list"
	"crypto/rand"
	"errors"
	"fmt"
)

type GameService struct {
}

func NewGameService() *GameService {
	return &GameService{}
}

var choices = []Choice{
	{ID: 1, Name: "Rock"},
	{ID: 2, Name: "Paper"},
	{ID: 3, Name: "Scissors"},
	{ID: 4, Name: "Lizard"},
	{ID: 5, Name: "Spock"},
}

var choicesMap = map[int]string{
	1: "Rock",
	2: "Paper",
	3: "Scissors",
	4: "Lizard",
	5: "Spock",
}

var superiority = map[int]map[int]bool{ //golang's ghetto map of int to set
	1: {4: true, 3: true}, //rock => lizard, scissors: see choicesMap for indices
	2: {1: true, 5: true}, //paper => rock, spock
	3: {2: true, 4: true}, //scissors => paper, lizard
	4: {2: true, 5: true}, //lizard => paper, spock
	5: {1: true, 3: true}, //spock => rock, scissors
}

var records = list.New()

var multimode = map[string]MultiMode{} //my fake, unexpiring cache.

func recordScore(record GameRecord) {
	records.PushFront(record)
	if records.Len() > 10 {
		i := records.Back()
		records.Remove(i)
	}
}

func (s *GameService) GetChoices() []Choice {
	return choices
}

func (s *GameService) GetRandomChoice() Choice {
	rand := GetRandomNumber()
	return choices[rand%5]
}

func getWinner(player1Option, player2Option int) string {
	winner := "Nobody" //tie
	if player1Option != player2Option {
		superlist := superiority[player1Option]
		if _, ok := superlist[player2Option]; ok {
			winner = "Player1"
		} else {
			winner = "Player2"
		}
	}
	return winner
}

func (s *GameService) Play(playRequest PlayRequest) (*PlayResponse, error) {
	rand := s.GetRandomChoice() //computer's choice
	if playRequest.Selection > 5 || playRequest.Selection < 1 {
		return nil, errors.New("Player choice must be between 1 and 5, inclusive")
	}

	pmap := map[string]string{
		"Nobody":  "Nobody",
		"Player1": "Computer",
		"Player2": "You",
	}

	winner := getWinner(rand.ID, playRequest.Selection)
	winner = pmap[winner]

	result := "tie"
	if winner == "Computer" { //computer
		result = "lose"
	} else if winner == "You" { //you
		result = "win"
	}

	//write to scoreboard.
	record := GameRecord{Player1: "Computer", Player2: "You", Player1Choice: rand.Name, Player2Choice: choicesMap[playRequest.Selection], Winner: winner}
	recordScore(record)
	return &PlayResponse{Result: result, PlayerSelection: playRequest.Selection, ComputerSelection: rand.ID}, nil
}

func (s *GameService) GetScoreboard() []GameRecord {
	//TODO: smells badly, rework
	var resolved []GameRecord
	for e := records.Front(); e != nil; e = e.Next() {
		resolved = append(resolved, e.Value.(GameRecord))
	}
	return resolved
}

func (s *GameService) ClearScoreboard() {
	records = list.New()
}

func (s *GameService) StartGame(challenge MultiModeStartRequest) (string, error) {
	key, err := getKey()
	if err != nil {
		return "", err
	}
	//got a valid key...
	start := MultiMode{Open: true, Player1: challenge.Player, Player1Choice: challenge.Selection}
	multimode[key] = start
	return key, nil
}

func (s *GameService) CompleteGame(complete MultiPlayerCompleteRequest) (MultiMode, error) {
	ongoing := multimode[complete.Key]
	if !ongoing.Open {
		return ongoing, nil
	}
	player2 := complete.Player
	if player2 == ongoing.Player1 {
		player2 = fmt.Sprintf("%s-The Usurper", ongoing.Player1) //fun enough huh?
	}
	ongoing.Player2 = player2
	ongoing.Player2Choice = complete.Selection

	pmap := map[string]string{
		"Nobody":  "Nobody",
		"Player1": ongoing.Player1,
		"Player2": ongoing.Player2,
	}

	winner := getWinner(ongoing.Player1Choice, ongoing.Player2Choice)
	winner = pmap[winner]
	ongoing.Winner = winner
	ongoing.Open = false
	multimode[complete.Key] = ongoing
	//
	record := GameRecord{Player1: ongoing.Player1, Player2: ongoing.Player2, Player1Choice: choicesMap[ongoing.Player1Choice], Player2Choice: choicesMap[ongoing.Player2Choice], Winner: winner}
	recordScore(record)
	return ongoing, nil
}

func (s *GameService) GetGameStatus(key string) MultiMode {
	return multimode[key]
}

func getKey() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	guid := fmt.Sprintf("%x", b[0:])
	return guid, nil
}
