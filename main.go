package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	appPort := os.Getenv("APP_PORT")
	router := gin.Default()

	router.GET("/healthz", HandleHealthz)

	router.Run(":" + appPort)
}

func HandleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}