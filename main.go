package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kehindeadewusi/rpssl-go-backend/game"
	"github.com/spf13/viper"
)

func main() {
	viperDefaults()

	router := setupRouter()
	log.Fatal(router.Run(fmt.Sprintf("%s:%v", viper.GetString("HOST"), viper.GetInt("PORT"))))
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	gameController := game.NewGameController()

	router.GET("/", gameController.GetChoices)
	router.GET("/choices", gameController.GetChoices)
	router.GET("/choice", gameController.GetRandomChoice)
	router.POST("/play", gameController.Play)
	router.GET("/scoreboard", gameController.GetScoreboard)
	router.POST("/clear-scoreboard", gameController.ClearScoreboard)
	router.POST("/start-multiplayer", gameController.StartMultiPlayer)
	router.POST("/complete-multiplayer", gameController.CompleteMultiPlayer)
	router.GET("/game-status/:key", gameController.GetGameStatus)

	http.Handle("/", router)

	return router
}

func viperDefaults() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RPSSL")
	//viper values here can be overriden with flags or env variables etc.
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8081")
}
