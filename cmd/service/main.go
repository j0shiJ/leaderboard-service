// cmd/service/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/j0shiJ/leaderboard-service/cmd/service/docs"
	"github.com/j0shiJ/leaderboard-service/internal/leaderboard"
	"github.com/j0shiJ/leaderboard-service/pkg/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// Import swaggerFiles from gin-swagger
)

// @title Leaderboard API
// @version 1.0
// @description This is a leaderboard service API.
// @host localhost:8080
// @BasePath /
func main() {
	db := database.InitDB()
	db.AutoMigrate(&leaderboard.User{})
	service := leaderboard.NewLeaderboardService(db)
	endpoints := leaderboard.MakeEndpoints(service)

	r := gin.Default()

	// Register the /swagger/*any route first
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register the other routes
	r.POST("/submit", gin.WrapH(httptransport.NewServer(
		endpoints.SubmitScoreEndpoint,
		leaderboard.DecodeSubmitScoreRequest,
		leaderboard.EncodeResponse,
	)))

	r.GET("/get_rank", gin.WrapH(httptransport.NewServer(
		endpoints.GetRankEndpoint,
		leaderboard.DecodeGetRankRequest,
		leaderboard.EncodeResponse,
	)))

	r.GET("/list_top_n", gin.WrapH(httptransport.NewServer(
		endpoints.ListTopNEndpoint,
		leaderboard.DecodeListTopNRequest,
		leaderboard.EncodeResponse,
	)))

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
