package main

import (
	"net/http"

	"github.com/allenliao0119/RestaurantForum/internal/config"
	"github.com/allenliao0119/RestaurantForum/internal/store/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn().
			Err(err).
			Msg("env file not found, using system env instead")
	}

	cfg, err := config.Environ()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to load config from env")
	}

	database.InitDB(cfg.Database)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, RestaurantForum!",
		})
	})

	r.Run(":8080")
}