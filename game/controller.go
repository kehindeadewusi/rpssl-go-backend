package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	service *GameService
}

func NewGameController() *GameController {
	return &GameController{service: NewGameService()}
}

func (ctr *GameController) GetChoices(c *gin.Context) {
	c.JSON(http.StatusOK, ctr.service.GetChoices())
}

func (ctr *GameController) GetRandomChoice(c *gin.Context) {
	c.JSON(http.StatusOK, ctr.service.GetRandomChoice())
}

func (ctr *GameController) Play(c *gin.Context) {
	var playRequest PlayRequest

	if err := c.ShouldBindJSON(&playRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	playResponse, err := ctr.service.Play(playRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, playResponse)
}

func (ctr *GameController) GetScoreboard(c *gin.Context) {
	c.JSON(http.StatusOK, ctr.service.GetScoreboard())
}

func (ctr *GameController) ClearScoreboard(c *gin.Context) {
	ctr.service.ClearScoreboard()
	c.JSON(http.StatusOK, gin.H{
		"message": "scoreboard reset successfully",
	})
}

func (ctr *GameController) StartMultiPlayer(c *gin.Context) {
	var multiPlayerRequest MultiModeStartRequest
	if err := c.ShouldBindJSON(&multiPlayerRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	key, err := ctr.service.StartGame(multiPlayerRequest)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"key": key})
}

func (ctr *GameController) CompleteMultiPlayer(c *gin.Context) {
	var completeGameRequest MultiPlayerCompleteRequest
	if err := c.ShouldBindJSON(&completeGameRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	multimode, err := ctr.service.CompleteGame(completeGameRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, multimode)
}

func (ctr *GameController) GetGameStatus(c *gin.Context) {
	key := c.Param("key")
	mode := ctr.service.GetGameStatus(key)
	if mode.Player1 == "" && mode.Player2 == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return //?
	}
	c.JSON(http.StatusOK, mode)
}
