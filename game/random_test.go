package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	result := GetRandomNumber()
	assert.GreaterOrEqual(t, result, 1, "Expecting number >= 1")
	assert.LessOrEqual(t, result, 100, "Expecting number <= 100")
}
