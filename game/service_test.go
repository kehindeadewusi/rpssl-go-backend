package game

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Rock = iota + 1
	Paper
	Scissors
	Lizard
	Spock
)

func TestPlayerSelectionLower(t *testing.T) {
	playRequest := PlayRequest{Selection: 0}
	_, err := NewGameService().Play(playRequest)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("Player choice must be between 1 and 5, inclusive"), err)
	}
}

func TestPlayerSelectionHigher(t *testing.T) {
	playRequest := PlayRequest{Selection: 6}
	_, err := NewGameService().Play(playRequest)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("Player choice must be between 1 and 5, inclusive"), err)
	}
}

func TestPlayerSelectionOK(t *testing.T) {
	playRequest := PlayRequest{Selection: 4}
	_, err := NewGameService().Play(playRequest)
	assert.Equal(t, err, nil)
}

func TestRockVsRock(t *testing.T) {
	winner := getWinner(Rock, Rock)
	assert.Equal(t, "Nobody", winner)
}

func TestRockVsSpock(t *testing.T) {
	winner := getWinner(Rock, Spock) //player1 rock, player2 spock
	assert.Equal(t, "Player2", winner)
}

func TestSpockVsRock(t *testing.T) {
	winner := getWinner(Spock, Rock) //player1 rock, player2 spock
	assert.Equal(t, "Player1", winner)
}
