package game

type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PlayRequest struct {
	Selection int `json:"player" binding:"required"`
}

type PlayResponse struct {
	Result            string `json:"results" binding:"required"`
	PlayerSelection   int    `json:"player" binding:"required"`
	ComputerSelection int    `json:"computer" binding:"required"`
}

type GameRecord struct {
	Player1       string `json:"player1"`
	Player2       string `json:"player2"`
	Player1Choice string `json:"player_1_choice"`
	Player2Choice string `json:"player_2_choice"`
	Winner        string `json:"winner"`
}

type MultiMode struct {
	Open          bool
	Player1       string
	Player2       string
	Player1Choice int
	Player2Choice int
	Winner        string
}

type MultiModeStartRequest struct {
	Player    string `json:"player" validate:"required"`
	Selection int    `json:"selection" validate:"required,min=1,max=5"`
}

type MultiPlayerCompleteRequest struct {
	Key       string `json:"key" binding:"required"`
	Player    string `json:"player" binding:"required"`
	Selection int    `json:"selection" binding:"required,min=1,max=5"`
}
