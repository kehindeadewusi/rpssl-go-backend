package docs

import "github.com/kehindeadewusi/rpssl-go-backend/game"

//swagger:route GET /choices game-tag choicesId
//Returns a list of player choices.
//responses:
//  200: choices

//A list of game choice.
//swagger:response choices
type choicesWrapper struct {
	//in:body
	Body []game.Choice
}

//swagger:route GET /choice game-tag choiceId
//Returns a random player choice.
//responses:
//  200: choice

//A random game choice.
//swagger:response choice
type choiceWrapper struct {
	//in:body
	Body game.Choice
}

//swagger:route POST /play game-tag playId
//Play a game, a player choice in the single-player mode.
//responses:
//  200: playResponse

//A player choice, computer choice is auto-generated on server
//swagger:response playResponse
type playerResponseWrapper struct {
	//in:body
	Body game.PlayResponse
}

// swagger:parameters playId
type playerRequestWrapper struct {
	// The value of the player selection has to be 1,2,3,4 or 5 to be valid.
	// in:body
	Body game.PlayRequest
}

//swagger:route GET /scoreboard game-tag scoreboardId
//Get the game scoreboard. Works for both single player and multi-player modes.
//responses:
//  200: scoreboardResponse

//A player choice, computer choice is auto-generated on server
//swagger:response scoreboardResponse
type gameRecordResponseWrapper struct {
	//in:body
	Body []game.GameRecord
}

//swagger:route POST /clear-scoreboard scoreboard-tag clearId
//Resets the game scoreboard. Affects single-player and multi-player mode.
//responses:
//  200:
//  500: clearResponseError

//An error message returned from the server, in the case of an API error from this endpoint.
//swagger:response clearResponseError
type clearResponseErrorWrapper struct {
	//in:body
	Body struct{ Message string }
}

//swagger:route POST /start-multiplayer multiplayer-tag startMultiId
//Starts a multi-player (2-player) mode. The initiator provides a name, plays a choice and gets back a correlation ID
//responses:
//  200: multiResponse

//A multi-player response, with a key for the other player to play their turn.
//swagger:response multiResponse
type multiResponseWrapper struct {
	//in:body
	Body struct {
		Key string `json:"key"`
	}
}

// swagger:parameters startMultiId
type multiRequestWrapper struct {
	// To start a multi-player game, the first player provides a name and their selection.
	// in:body
	Body game.MultiModeStartRequest
}

//swagger:route POST /complete-multiplayer multiplayer-tag completeMultiId
//Second player move in a multi-player (2-player) mode. It finds an existing game with a key and responds.
//responses:
//  200: completeMultiResponse

//A response to the second player move in a multi-player game.
//swagger:response completeMultiResponse
type completeMultiResponseWrapper struct {
	//in:body
	Body game.MultiMode
}

// swagger:parameters completeMultiId
type completeMultiRequestWrapper struct {
	// To play a move in a multi-player game, the second player submits the game key, their name and choice.
	// in:body
	Body game.MultiPlayerCompleteRequest
}

//swagger:route GET /game-status/{key} multiplayer-tag gameStatusId
//Gets the status of a multi-player game, identified by the key (a correlation ID)
//responses:
//  200: scoreboardResponse

//A player choice, computer choice is auto-generated on server
//swagger:response scoreboardResponse
type gameStatusResponseWrapper struct {
	//in:body
	Body game.MultiMode
}
