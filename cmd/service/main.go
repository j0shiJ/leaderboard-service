// cmd/service/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/j0shiJ/leaderboard-service/cmd/service/docs"
	"github.com/j0shiJ/leaderboard-service/internal/leaderboard"
	"github.com/j0shiJ/leaderboard-service/pkg/database"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// Import swaggerFiles from gin-swagger
)

// title Leaderboard API
// version 1.0
// description This is a leaderboard service API.
// host localhost:8080
// BasePath /
func main() {
	db := database.InitDB()
	db.AutoMigrate(&leaderboard.User{})
	service := leaderboard.NewLeaderboardService(db)
	endpoints := leaderboard.MakeEndpoints(service)
	handler := leaderboard.MakeHTTPHandler(endpoints)

	r := gin.Default()
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Any("/*proxyPath", gin.WrapH(handler))

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
